package client

import (
	"DHT/internal/api"
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
)

// Client defines a API client connecting to a API server representing the Chord network.
type Client struct {
	Address string
	conn    net.Conn
	reader  *bufio.Reader
}

// NewClient creates a client connecting to the given API server address.
func NewClient(address string) *Client {
	c := &Client{Address: address}
	var err error
	c.conn, err = net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Client error:", err)
		return nil
	}
	c.reader = bufio.NewReader(c.conn)
	return c
}

// sendGetMessage sends a GET message to the server.
func (c *Client) sendGetMessage(key []byte) error {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint16((32+256)/8))  // size
	binary.Write(buf, binary.BigEndian, uint16(api.DHT_GET)) // DHT GET
	if len(key) < 32 {                                       // key
		buf.Write(key)
		for i := 0; i < 32-len(key); i++ {
			buf.WriteByte(0)
		}
	} else {
		buf.Write(key[:32])
	}
	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// sendPutMessage sends a PUT message to the server.
func (c *Client) sendPutMessage(key []byte, value []byte, ttl uint16, replication uint8) error {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, uint16((64+256)/8+len(value))) // size
	binary.Write(buf, binary.BigEndian, uint16(api.DHT_PUT))           // DHT PUT
	binary.Write(buf, binary.BigEndian, uint16(ttl))                   // ttl
	binary.Write(buf, binary.BigEndian, uint8(replication))            // replication
	binary.Write(buf, binary.BigEndian, uint8(0))                      // reserved
	if len(key) < 32 {                                                 // key
		buf.Write(key)
		for i := 0; i < 32-len(key); i++ {
			buf.WriteByte(0)
		}
	} else {
		buf.Write(key[:32])
	}
	buf.Write(value) // value

	_, err := c.conn.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

// receiveMessage receives a DHT_SUCCESS or DHT_FAILURE message from the server.
func (c *Client) receiveMessage() ([]byte, bool, error) {
	data := make([]byte, 4+32)
	if n, err := c.reader.Read(data); err != nil || n != 4+32 {
		return nil, false, errors.New("message length error")
	}
	size := binary.BigEndian.Uint16(data[0:2])
	msgType := binary.BigEndian.Uint16(data[2:4])
	if api.MsgType(msgType) == api.DHT_FAILURE {
		return nil, false, nil
	}
	if api.MsgType(msgType) != api.DHT_SUCCESS {
		return nil, false, errors.New("message error")
	}
	data = make([]byte, size-4-32)
	if n, err := c.reader.Read(data); err != nil || n != len(data) {
		return nil, false, errors.New("message length error")
	}
	return data, true, nil
}

// Get retrieves the value for the key from the server.
func (c *Client) Get(key []byte) ([]byte, bool, error) {
	if err := c.sendGetMessage(key); err != nil {
		return nil, false, err
	}
	return c.receiveMessage()
}

// Put ask the server to store the key/value pair to the Chord network.
func (c *Client) Put(key []byte, value []byte, ttl uint16, replication uint8) {
	err := c.sendPutMessage(key, value, ttl, replication)
	if err != nil {
		return
	}
}

// Close closes the connection.
func (c *Client) Close() {
	c.conn.Close()
	c.conn = nil
}
