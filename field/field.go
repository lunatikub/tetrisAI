package field

import (
	tetromino "github.com/lunatikub/tetrisAI/tetromino"
)

const (
	fieldHeight = 20
	fieldWidth  = 10
)

// Field blocks
type Field struct {
	Blocks [fieldHeight][fieldWidth]int
	Height [fieldWidth]int
}

func (f *Field) GetHeight(p *tetromino.Piece, x int) int {
	h := 0
	for i := 0; i < p.Width; i++ {
		if f.Height[i]+1 > p.HeightCol[i] {
			hCol := f.Height[i] - p.HeightCol[i] + 1
			if h < hCol {
				h = hCol
			}
		}
	}
	return h
}

// func (f *field) push(tetro int, x int, rot int) {

// }
