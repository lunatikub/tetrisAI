package tetromino

// Tetromino piece
type piece struct {
	cells     [][]int // Tetromino cells
	heightCol []int   // Height by column
	height    int
	width     int
}

// tetromino is a set of piece rotations
type tetromino struct {
	name   string
	pieces []piece // tetromino piece rotations
}

func (p *piece) isFilled(y int, x int) bool {
	return p.cells[y][x] == 1
}
