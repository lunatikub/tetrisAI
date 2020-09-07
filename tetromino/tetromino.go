package tetromino

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
type Piece struct {
	Blocks    [][]int
	HeightCol []int
	Height    int
	Width     int
}

// Definition is a set of piece rotations
type Definition struct {
	Name   string
	Pieces []Piece // tetromino piece rotations
}
