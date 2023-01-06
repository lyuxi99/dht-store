package chord

import (
	"DHT/internal/chord/proto"
	"DHT/internal/logger"
	"DHT/internal/storage"
	"DHT/internal/utils"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

// P2pServer defines a P2P server handling any requests from other P2P servers.
type P2pServer struct {
	RpcServer  *ChordRpcServer // the underlying rpc server of type ChordRpcServer
	RpcService *grpc.Server    // the current running rpc service of the underlying rpc server
	lis        net.Listener    // the net.Listener which the underlying RpcServer should listen on
	stopped    bool            // whether the P2pServer has stopped
	wg         *sync.WaitGroup // used for graceful shutdown
}

// loadServerTLSCredentials loads TLS Credentials from the given cert and key for server
func loadServerTLSCredentials(caCertFile string, serverCert, serverKey string) (credentials.TransportCredentials, error) {
	b, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}

	cert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, err
	}
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    cp,
	}), nil
}

// loadClientTLSCredentials loads TLS Credentials from the given cert and key for client
func loadClientTLSCredentials(caCertFile string, serverCert, serverKey string) (credentials.TransportCredentials, error) {
	b, err := ioutil.ReadFile(caCertFile)
	if err != nil {
		return nil, err
	}
	cp := x509.NewCertPool()
	if !cp.AppendCertsFromPEM(b) {
		return nil, fmt.Errorf("credentials: failed to append certificates")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(serverCert, serverKey)
	if err != nil {
		return nil, err
	}

	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      cp,
	}), nil
}

// NewP2pServer creates a new P2P server listening on the given address.
func NewP2pServer(storage *storage.Storage, address string, caCert, serverCert, serverKey string) *P2pServer {
	serverCreds, err := loadServerTLSCredentials(caCert, serverCert, serverKey)
	if err != nil {
		log.Fatal(err)
	}
	clientCreds, err := loadClientTLSCredentials(caCert, serverCert, serverKey)
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(serverCreds))
	server := NewChordServer(storage, address, clientCreds)
	proto.RegisterChordServer(s, server)
	p2pServer := &P2pServer{
		RpcServer:  server,
		RpcService: s,
		lis:        lis,
		stopped:    false,
		wg:         &sync.WaitGroup{},
	}
	return p2pServer
}

// Serve runs the chord service, and joins the existing Chord network if bootstrapper is specified, otherwise create a new Chord network.
func (s *P2pServer) Serve(bootstrapper string) error {
	s.wg.Add(1)
	go func(s *P2pServer) {
		time.Sleep(time.Second)
		if len(strings.TrimSpace(bootstrapper)) > 0 {
			// Join the existing chord
			if err := s.RpcServer.Join(context.Background(), NewNode(bootstrapper)); err != nil {
				log.Fatal("server.Join error", err)
			}
			fmt.Println("server.Join ok!")
		}

		// Stabilize
		lastInfo := ""
		for {
			if s.stopped {
				break
			}
			err := s.RpcServer.CheckPredecessorAndSuccessor(context.Background())
			if err != nil {
				if s.stopped {
					break
				}
				logger.Logger.Warnw("server.Stabilize error", "err", err)
			}
			if s.stopped {
				break
			}
			err = s.RpcServer.Stabilize(context.Background())
			if err != nil {
				if s.stopped {
					break
				}
				logger.Logger.Warnw("server.Stabilize error", "err", err)
			}
			if s.stopped {
				break
			}
			err = s.RpcServer.FixFingers(context.Background())
			if err != nil {
				if s.stopped {
					break
				}
				logger.Logger.Warnw("server.FixFingers error", "err", err)
			}
			if s.stopped {
				break
			}
			newInfo := s.RpcServer.GetInfoString()
			if lastInfo != newInfo {
				lastInfo = newInfo
				fmt.Println(newInfo)
			}
			if s.stopped {
				break
			}
			//time.Sleep(time.Second)
			time.Sleep(time.Millisecond * 10)
		}
		fmt.Println("P2p server stopped!")
		s.wg.Done()
	}(s)
	if err := s.RpcService.Serve(s.lis); err != nil {
		return err
	}

	return nil
}

// Stop stops the P2pServer gracefully.
func (s *P2pServer) Stop() {
	s.stopped = true
	s.RpcService.GracefulStop()
	s.wg.Wait()
}

// ChordRpcServer defines a Chord server running the Chord algorithm.
type ChordRpcServer struct {
	proto.UnimplementedChordServer
	storage       *storage.Storage
	Self          *Node
	Finger        []*Node
	Predecessor   *Node
	successorList []*Node
	mutex         sync.Mutex
	ClientCreds   credentials.TransportCredentials
}

// NewChordServer creates a new Chord server with the given underlying storage.Storage, listening on the given address.
func NewChordServer(storage *storage.Storage, addr string, clientCreds credentials.TransportCredentials) *ChordRpcServer {
	s := &ChordRpcServer{
		Self:        NewNode(addr),
		Finger:      make([]*Node, M),
		Predecessor: nil,
		storage:     storage,
		ClientCreds: clientCreds,
	}
	//s.Predecessor = s.Self
	for i := 0; i < M; i++ {
		s.Finger[i] = s.Self
	}
	return s
}

// successor returns s.Finger[0], which is our successor.
func (s *ChordRpcServer) successor() *Node {
	if len(s.Finger) > 0 {
		return s.Finger[0]
	}
	return nil
}

// closestPrecedingFinger return the closest preceding finger of the given id.
func (s *ChordRpcServer) closestPrecedingFinger(id []byte) *Node {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if len(s.Finger) == 0 {
		return nil
	}
	for i := len(s.Finger) - 1; i >= 0; i-- {
		if utils.IsInRangeExclude(s.Finger[i].Id, s.Self.Id, id) {
			return s.Finger[i]
		}
	}
	return s.Self
}

// GetInfoString get the information of server as a string.
func (s *ChordRpcServer) GetInfoString() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Predecessor: %v\n", s.Predecessor.ToString()))
	sb.WriteString(fmt.Sprintf("Self: %v\n", s.Self.ToString()))
	//fmt.Println("Predecessor:", s.Predecessor.ToString())
	sb.WriteString("fingers:")
	//var indices, ids, ips []string
	for i, node := range s.Finger {
		if i > 2 {
			break
		}
		//indices = append(indices, strconv.Itoa(i))
		sb.WriteString(fmt.Sprintf("(%v: ", i))
		if node == nil {
			sb.WriteString("nil) ")
			//ids = append(ids, "nil")
			//ips = append(ips, "nil")
		} else {
			//sb.WriteString(hex.EncodeToString(node.Id)[:4] + ":")
			sb.WriteString(hex.EncodeToString(node.Id)[:] + ":")
			sb.WriteString(node.Addr[len(node.Addr)-4:] + ") ")
			//ids = append(ids, hex.EncodeToString(node.Id)[:4])
			//ids = append(ids, node.Addr[len(node.Addr)-4:])
		}
		//table[0] = append(table[0])
		//sb.WriteString(fmt.Sprintf("Finger+2**%d: %s\n", i, node.ToString()))
		//fmt.Printf("Finger+2**%d: %s\n", i, node.ToString())
	}
	sb.WriteString(" ...\n")
	sb.WriteString("successorList: ")
	for _, node := range s.successorList {
		sb.WriteString(node.ToString() + ", ")
	}
	sb.WriteString("\n")

	//buf := bytes.NewBufferString("")
	//w := tabwriter.NewWriter(buf, 1, 1, 1, ' ', 0)
	//fmt.Fprintln(w, strings.Join(indices, "\t")+"\t")
	//fmt.Fprintln(w, strings.Join(ids, "\t")+"\t")
	//fmt.Fprintln(w, strings.Join(ips, "\t")+"\t")
	//w.Flush()

	//sb.WriteString(buf.String())
	return sb.String()
}

// logFunc logs the request and the response of the function call.
func logFunc(name string, req interface{}, resp interface{}, err error) {
	if err == nil {
		logger.Logger.Debugw(name, "req", req, "resp", resp, "err", err)
	} else {
		logger.Logger.Warnw(name, "req", req, "resp", resp, "err", err)
	}
}

// FindSuccessor asks us to find the given id's successor.
func (s *ChordRpcServer) FindSuccessor(ctx context.Context, id *proto.Id) (resp *proto.Node, err error) {
	defer logFunc("s.FindSuccessor", id, resp, err)
	// if successor(s) is the successor of the id
	if utils.IsInRange(id.Id, s.Self.Id, s.successor().Id) {
		return s.successor().ToProtoNode(), nil
	}
	// else
	nn := s.closestPrecedingFinger(id.Id)
	if nn == nil {
		return nil, status.Errorf(codes.Unknown, "closestPrecedingFinger() is nil")
	}
	nnc, err := nn.GetClient(s.ClientCreds)
	if err != nil {
		return nil, err
	}
	res, err := nnc.FindSuccessor(ctx, id)
	if err != nil {
		return nil, err
	}
	return &proto.Node{Id: res.Id, Addr: res.Addr}, nil
}

// GetPredecessor returns our Predecessor.
func (s *ChordRpcServer) GetPredecessor(ctx context.Context, in *proto.Void) (resp *proto.Node, err error) {
	defer logFunc("s.GetPredecessor", in, resp, err)
	if s.Predecessor == nil {
		return nil, status.Errorf(codes.Unknown, "s.Predecessor is nil")
	}
	return s.Predecessor.ToProtoNode(), nil
}

/*=====================================================
                    Stabilization
=====================================================*/

// Join the chord system through a broker.
func (s *ChordRpcServer) Join(ctx context.Context, bootstrapper *Node) (err error) {
	defer logFunc("s.Join", bootstrapper, nil, err)
	s.Predecessor = nil

	// ask bootstrapper for our successor
	c, err := bootstrapper.GetClient(s.ClientCreds)
	if err != nil {
		return err
	}
	suc, err := c.FindSuccessor(ctx, &proto.Id{Id: s.Self.Id})
	if err != nil {
		return err
	}
	s.Finger[0] = NewNodeFromProtoNode(suc)
	return nil
}

// CheckPredecessorAndSuccessor checks whether the Predecessor and the successor are still alive.
func (s *ChordRpcServer) CheckPredecessorAndSuccessor(ctx context.Context) error {
	c, err := s.successor().GetClient(s.ClientCreds)
	if err == nil {
		_, err = c.Ping(ctx, &proto.Void{})
	}
	s.mutex.Lock()
	if err != nil {
		if len(s.successorList) >= 2 {
			s.Finger[0] = s.successorList[1]
			s.successorList = s.successorList[1:]
		} else {
			s.Finger[0] = s.Self
		}
	}
	s.mutex.Unlock()
	if s.Predecessor == nil {
		return nil
	}
	c, err = s.Predecessor.GetClient(s.ClientCreds)
	if err == nil {
		_, err = c.Ping(ctx, &proto.Void{})
	}
	if err != nil {
		s.Predecessor = nil
	}
	return nil
}

// Stabilize periodically verify our immediate successor, and tell the successor about us.
func (s *ChordRpcServer) Stabilize(ctx context.Context) (err error) {
	defer logFunc("s.Stabilize", nil, nil, err)
	c, err := s.successor().GetClient(s.ClientCreds)
	if err != nil {
		return err
	}
	x, err := c.GetPredecessor(ctx, &proto.Void{})
	if err == nil && utils.IsInRangeExclude(x.Id, s.Self.Id, s.successor().Id) {
		if utils.CheckIdentity(x.Id, x.Addr) {
			s.Finger[0] = NewNodeFromProtoNode(x)
		} else {
			logger.Logger.Warnw("Stabilize: identity check error", "node", x)
		}
	}
	cSuc, err := s.successor().GetClient(s.ClientCreds)
	if err != nil {
		return err
	}
	resp, err := cSuc.Notify(ctx, s.Self.ToProtoNode())
	if err != nil {
		return err
	}
	s.mutex.Lock()
	s.successorList = []*Node{}
	for _, node := range resp.GetNodes() {
		s.successorList = append(s.successorList, NewNodeFromProtoNode(node))
	}
	s.mutex.Unlock()
	return nil
}

// Notify lets us think `nn` might be our Predecessor.
func (s *ChordRpcServer) Notify(ctx context.Context, nn *proto.Node) (resp *proto.SuccessorList, err error) {
	defer logFunc("s.Notify", nn, resp, err)
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// verify node, to avoid ID mapping attacking
	if !utils.CheckIdentity(nn.Id, nn.Addr) {
		logger.Logger.Warnw("Notify: identity check error", "node", nn)
		return nil, status.Error(codes.PermissionDenied, "identity check error")
	}

	if s.Predecessor == nil || utils.IsInRangeExclude(nn.Id, s.Predecessor.Id, s.Self.Id) {
		s.Predecessor = NewNodeFromProtoNode(nn)
		// no need to transfer data
	}
	resp = &proto.SuccessorList{Nodes: []*proto.Node{s.Self.ToProtoNode()}}
	for _, node := range s.successorList {
		if len(resp.Nodes) >= NUM_SUCCESSORS_IN_LIST {
			break
		}
		resp.Nodes = append(resp.Nodes, node.ToProtoNode())
	}
	return resp, nil
}

// FixFingers refreshes a random Finger table entry, should be called periodically.
func (s *ChordRpcServer) FixFingers(ctx context.Context) (err error) {
	defer logFunc("s.FixFingers", nil, nil, err)
	i := rand.Intn(M-1) + 1 // random integer number in [1,M)
	id := utils.AddBytesPower2(s.Self.Id, i)
	node, err := s.FindSuccessor(ctx, &proto.Id{Id: id})
	if err != nil {
		return err
	}
	s.Finger[i] = NewNodeFromProtoNode(node)
	return nil
}

// Ping asks us to respond with an empty message, used to keep alive
func (s *ChordRpcServer) Ping(ctx context.Context, in *proto.Void) (*proto.Void, error) {
	return &proto.Void{}, nil
}

/*=====================================================
                 Storage Operations
=====================================================*/

// Put asks us to put the key/value pair to our storage, then forwards the request to our successor if needed.
func (s *ChordRpcServer) Put(ctx context.Context, req *proto.PutReq) (resp *proto.Void, err error) {
	defer logFunc("s.Put", req, resp, err)
	if req.GetInitiatorAddr() == "" {
		req.InitiatorAddr = s.Self.Addr
	} else if s.Self.Addr == req.GetInitiatorAddr() {
		return &proto.Void{}, nil
	}
	ttl := time.UnixMilli(req.Expire).Sub(time.Now())
	if ttl.Milliseconds() > 0 {
		s.storage.Put(req.Key, req.Value, ttl)
	}
	// forward the request to successor
	if req.Replication <= 1 {
		return &proto.Void{}, nil
	}
	c, err := s.Finger[0].GetClient(s.ClientCreds)
	if err != nil {
		return nil, err
	}
	_, err = c.Put(ctx, &proto.PutReq{
		Key:           req.Key,
		Value:         req.Value,
		Expire:        req.Expire,
		InitiatorAddr: req.InitiatorAddr,
		Replication:   req.Replication - 1,
	})
	if err != nil {
		return nil, err
	}
	return &proto.Void{}, nil
}

// Get asks us to get the value for the given key from our storage.
func (s *ChordRpcServer) Get(ctx context.Context, req *proto.GetReq) (resp *proto.GetResp, err error) {
	defer logFunc("s.Put", req, resp, err)
	val, ok := s.storage.Get(req.GetKey())
	return &proto.GetResp{Value: val, Ok: ok}, nil
}
