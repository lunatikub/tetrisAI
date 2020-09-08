package player

//  +---------+---------+---------+
//  | I: XXXX | L: XXX  | J: XXX  |
//  |         |      X  |    X    |
//  +---------+---------+---------+
//  | O:  XX  | S:  XX  | T:  X   |
//  |     XX  |    XX   |    XXX  |
//  +---------+---------+---------+
//  | Z: XX   |
//  |     XX  |
//  +---------+
//  */
const (
	I = iota
	J
	L
	O
	S
	T
	Z
)

// Piece rotation
type piece struct {
	blocks    [][]int
	holes     []int
	heightCol []int
	height    int
	width     int
}

// A tetromino is a set of piece rotations
type tetromino struct {
	name   string
	pieces []piece // tetromino piece rotations
}
