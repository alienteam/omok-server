package core

import (
	"github.com/go-redis/redis"
)

// MessageBroker ...
type MessageBroker struct {
	Redis *redis.Client
}

// Publish ...
func (r *MessageBroker) Publish() {}

// Subscribe ...
func (r *MessageBroker) Subscribe() {}
