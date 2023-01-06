package chord

import (
	"DHT/internal/chord/proto"
	"DHT/internal/utils"
	"bytes"
	"encoding/hex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Node defines a chord node, which can be used for communication
type Node struct {
	Id         []byte            // the id of this node
	Addr       string            // the address of format ip:p2p_port
	clientConn *grpc.ClientConn  // the grpc.ClientConn connecting to this node
	client     proto.ChordClient // the proto.ChordClient, used for sending any requests
}

// GetClient returns the proto.ChordClient for sending any requests to this node
func (p *Node) GetClient(clientCreds credentials.TransportCredentials) (proto.ChordClient, error) {
	if p.client != nil {
		return p.client, nil
	}
	conn, err := grpc.Dial(p.Addr, grpc.WithTransportCredentials(clientCreds))
	if err != nil {
		return nil, err
	}
	c := proto.NewChordClient(conn)
	p.client = c
	p.clientConn = conn

	//runtime.SetFinalizer(p, p.Close)
	return p.client, nil
}

// Close closes the connection to this node
func (p *Node) Close() {
	if p.clientConn != nil {
		p.clientConn.Close()
		p.clientConn = nil
		p.client = nil
	}
}

// ToString returns a string representing the metadata of this node
func (p *Node) ToString() string {
	if p == nil {
		return "nil"
	}
	sb := bytes.NewBufferString("(")
	sb.WriteString(hex.EncodeToString(p.Id))
	sb.WriteString(", ")
	sb.WriteString(p.Addr)
	sb.WriteString(")")
	return sb.String()
}

// ToProtoNode creates a proto.Node of this node
func (p *Node) ToProtoNode() *proto.Node {
	return &proto.Node{
		Id:   p.Id,
		Addr: p.Addr,
	}
}

// NewNodeFromProtoNode creates a Node from the given proto.Node
func NewNodeFromProtoNode(p *proto.Node) *Node {
	return &Node{
		Id:   p.Id,
		Addr: p.Addr,
	}
}

// NewNode creates a Node from the given address
func NewNode(addr string) *Node {
	return &Node{
		Id:   utils.SHA1([]byte(addr)),
		Addr: addr,
	}
}
