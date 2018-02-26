package game

// Stone represents a omok stone.
type Stone byte

const (
	// NoStone is Empty Stone.
	NoStone = 0
	// BlackStone is black color stone.
	BlackStone = 1
	// WhiteStone is white color stone.
	WhiteStone = 2
)

// Omok represents omok board and logic.
type Omok struct {
	board [10][10]Stone
	turn  byte
}

// Place place a stone to the (x,y) position.
// And return true, if the stone wins, else return false.
func (o *Omok) Place(x, y int, stone Stone) bool {
	o.board[y][x] = stone
	return o.Win(x, y, stone)
}

// Win checks the stone wins.
func (o *Omok) Win(x, y int, stone Stone) bool {
	return false
}
