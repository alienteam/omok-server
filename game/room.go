package game

// Room represents a battle room.
type Room struct {
	id       int
	user     [2]User
	observer []User
	omok     Omok
}

func (r *Room) AddUser(u *User) {
}

func (r *Room) RemoveUser(u *User) {
}

// NewRoom creates a Room.
func NewRoom(user User) *Room {
	r := &Room{}
	r.user[0] = user
	return r
}
