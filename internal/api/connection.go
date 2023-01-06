package api

import (
	"DHT/internal/logger"
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
)

// Connection defines a connection to a client.
type Connection struct {
	s        *ApiServer    // the ApiServer who creates this Connection
	conn     net.Conn      // the net.Conn which the Connection is listening to
	reader   *bufio.Reader // the bufio.Reader for the conn
	sendLock sync.Mutex    // the sync.Mutex for sending any messages to the client
}

// NewConnection creates a connection to a client of the given net.Conn object.
func NewConnection(s *ApiServer, conn net.Conn) *Connection {
	return &Connection{s: s, conn: conn, reader: bufio.NewReader(conn)}
}

// threadReceiveMsg listens to a client and handle incoming requests.
func (p *Connection) threadReceiveMsg() {
	defer func() {
		logger.Logger.Infow("connection closed!", "addr", p.conn.RemoteAddr())
		p.conn.Close()
		p.conn = nil
	}()

	for true {
		msgType, msgBody := p.readMessage()
		if msgBody == nil {
			break
		}
		logger.Logger.Infow("readMessage", "msgType", msgType, "msgBody", string(msgBody), "addr", p.conn.RemoteAddr())
		go p.handleMessage(msgType, msgBody)
	}
}

// handleMessage handles an incoming request, and returns the response message, if any.
func (p *Connection) handleMessage(msgType MsgType, msgBody []byte) {
	defer func() {
		if err := recover(); err != nil {
			logger.Logger.Errorw("panic when handleMessage", "err", err, "msgType", msgType, "msgBody", string(msgBody))
		}
	}()
	respMsgType, respMsgBody := p.s.ProcessMessage(msgType, msgBody)
	if respMsgType != 0 {
		logger.Logger.Infow("sendMessage", "msgType", respMsgType, "msgBody", string(respMsgBody), "addr", p.conn.RemoteAddr())
		if err := p.sendMessage(respMsgType, respMsgBody); err != nil {
			logger.Logger.Warnw("sendMessage error", "err", err, "addr", p.conn.RemoteAddr())
		}
	}
}

// readMessage reads an incoming request, return its message type in MsgType and message body in []byte.
func (p *Connection) readMessage() (MsgType, []byte) {
	data := make([]byte, 4)
	if n, err := p.reader.Read(data); err != nil || n != 4 {
		if err == io.EOF {
			return 0, nil
		}
		logger.Logger.Warnw("readMessage header error", "n", n, "err", err)
		return 0, nil
	}
	size := binary.BigEndian.Uint16(data[0:2])
	msgType := binary.BigEndian.Uint16(data[2:4])

	data = make([]byte, size-4)
	if n, err := p.reader.Read(data); err != nil || n != int(size)-4 {
		logger.Logger.Warnw("readMessage body error", "n", n, "err", err)
		return 0, nil
	}
	return MsgType(msgType), data
}

// sendMessage sends a response message defined by message type and message body, to the client.
func (p *Connection) sendMessage(msgType MsgType, msgBody []byte) error {
	p.sendLock.Lock()
	defer p.sendLock.Unlock()
	buf := bytes.NewBuffer([]byte{})
	size := 4 + len(msgBody)
	if err := binary.Write(buf, binary.BigEndian, uint16(size)); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, uint16(msgType)); err != nil {
		return err
	}
	buf.Write(msgBody)
	if p.conn == nil {
		return errors.New("p.conn is nil")
	}
	_, err := p.conn.Write(buf.Bytes())
	return err
}
