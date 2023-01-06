package api

// MsgType is a type defining the message type of any messages when communicating with clients.
type MsgType uint16

// Here defines some MsgType constants for all types of messages.
const (
	DHT_PUT     MsgType = 650
	DHT_GET     MsgType = 651
	DHT_SUCCESS MsgType = 652
	DHT_FAILURE MsgType = 653
)
