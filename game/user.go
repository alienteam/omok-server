package game

import "github.com/alienteam/omok-server/core"

// User represents a user.
type User struct {
	conn *core.Connection
	uid  int
	name string
}
