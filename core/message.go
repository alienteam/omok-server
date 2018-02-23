package core

import "encoding/json"

// Message is a unit of data
type Message interface{}

// MessageHandler is message interface
type MessageHandler interface {
	Encode(m Message) []byte
	Decode(v []byte) Message
}

// JsonMessageHandler is json type of MessageHandler
type JsonMessageHandler struct{}

// Encode encodes json type.
func (h *JsonMessageHandler) Encode(v Message) []byte {
	msg, _ := json.Marshal(v)
	return msg
}

// Decode decodes json type.
func (h *JsonMessageHandler) Decode(v []byte) Message {
	var t interface{}
	json.Unmarshal(v, &t)
	return t
}

type StringMessageHandler struct{}

// Encode encodes string type.
func (h *StringMessageHandler) Encode(v Message) []byte {
	return []byte(v.(string))
}

// Decode decodes string type.
func (h *StringMessageHandler) Decode(v []byte) Message {
	return string(v)
}
