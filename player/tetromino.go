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
	nrTetromino
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

func getTetromino(tetromino int) *tetromino {
	switch tetromino {
	case I:
		return &tetrominoI
	case J:
		return &tetrominoJ
	case L:
		return &tetrominoL
	case O:
		return &tetrominoO
	case S:
		return &tetrominoS
	case T:
		return &tetrominoT
	case Z:
		return &tetrominoZ
	}
	panic("Not available tetromino")
}
