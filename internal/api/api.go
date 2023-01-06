package api

import (
	"DHT/internal/chord"
	"DHT/internal/chord/proto"
	"DHT/internal/logger"
	"DHT/internal/utils"
	"context"
	"encoding/binary"
	"time"
)

// Put the key/value pair into the storage expiring in `ttl` seconds,
// and the pair should be replicated for `replication` times.
func (s *ApiServer) Put(key []byte, value []byte, ttl uint16, replication uint8) {
	logger.Logger.Infow("api.Put", "key", string(key), "value", string(value), "ttl", ttl, "replication", replication)
	var err error
	defer func() {
		if err != nil {
			logger.Logger.Infow("api.Put error", "err", err)
		}
	}()
	// find successor
	// initiate put request
	expire := time.Now().Add(time.Second * time.Duration(ttl)).UnixMilli()
	respNode, err := s.p2pServer.RpcServer.FindSuccessor(context.Background(), &proto.Id{Id: utils.SHA1(key)})
	if err != nil {
		return
	}
	node := chord.NewNodeFromProtoNode(respNode)
	c, err := node.GetClient(s.p2pServer.RpcServer.ClientCreds)
	if err != nil {
		return
	}
	defer node.Close()
	req := &proto.PutReq{
		Key:           key,
		Value:         value,
		Expire:        expire,
		InitiatorAddr: "",
		Replication:   int32(replication),
	}
	resp, err := c.Put(context.Background(), req)
	logger.Logger.Infow("api.Put over", "node", node, "req", req, "resp", resp, "err", err)
}

// Get finds the value for the given key, if any.
func (s *ApiServer) Get(key []byte) ([]byte, bool) {
	logger.Logger.Infow("api.Get", "key", string(key))
	var err error
	defer func() {
		if err != nil {
			logger.Logger.Infow("api.Get error", "err", err)
		}
	}()
	// find successor
	// initiate put
	respNode, err := s.p2pServer.RpcServer.FindSuccessor(context.Background(), &proto.Id{Id: utils.SHA1(key)})
	if err != nil {
		return nil, false
	}
	node := chord.NewNodeFromProtoNode(respNode)
	c, err := node.GetClient(s.p2pServer.RpcServer.ClientCreds)
	if err != nil {
		return nil, false
	}
	defer node.Close()
	req := &proto.GetReq{Key: key}
	resp, err := c.Get(context.Background(), req)
	if err != nil {
		return nil, false
	}
	logger.Logger.Infow("api.Get over", "node", node, "req", req, "resp", resp, "err", err)
	return resp.GetValue(), resp.GetOk()
}

// ProcessMessage processes the given message, and return the response message, otherwise return 0, nil.
func (s *ApiServer) ProcessMessage(msgType MsgType, msgBody []byte) (MsgType, []byte) {
	switch msgType {
	case DHT_PUT:
		TTL := binary.BigEndian.Uint16(msgBody[0:2])
		replication := msgBody[2]
		key := msgBody[4:36]
		value := msgBody[36:]
		s.Put(key, value, TTL, replication)
		return 0, nil
	case DHT_GET:
		key := msgBody[0:32]
		if value, ok := s.Get(key); ok {
			data := key
			data = append(data, value...)
			return DHT_SUCCESS, data
		}
		return DHT_FAILURE, key
	default:
		return 0, nil
	}
}
