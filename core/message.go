package core

import "encoding/json"

// Message is unit of data.
type Message interface {
	Encode(v interface{}) interface{}
	Decode(v interface{}) interface{}
}

// JsonMessage is json type of message
type JsonMessage struct {
}

// Encode encodes json type.
func (m *JsonMessage) Encode(data map[string]interface{}) interface{} {
	msg, _ := json.Marshal(m)
	return msg
}

// Decode decodes json type.
func (m *JsonMessage) Decode(data []byte) map[string]interface{} {
	var t interface{}
	json.Unmarshal(data, &t)
	return t.(map[string]interface{})
}
