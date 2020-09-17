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
	nrTetrimino
)

// Piece rotation
type piece struct {
	blocks    [][]int
	holes     []int
	heightCol []int
	blockRow  []int
	height    int
	width     int
}

// A tetrimino is a set of piece rotations
type tetrimino struct {
	name   string
	pieces []piece // tetrimino piece rotations
}

func getTetrimino(tetrimino int) *tetrimino {
	switch tetrimino {
	case I:
		return &tetriminoI
	case J:
		return &tetriminoJ
	case L:
		return &tetriminoL
	case O:
		return &tetriminoO
	case S:
		return &tetriminoS
	case T:
		return &tetriminoT
	case Z:
		return &tetriminoZ
	}
	panic("Not available tetrimino")
}
