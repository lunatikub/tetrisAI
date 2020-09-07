package tetromino

import (
	"testing"
)

func (t *tetromino) TestTetromino(test *testing.T) {
	for r, p := range t.pieces {

		if p.width != len(p.cells[0]) || p.height != len(p.cells) {
			test.Errorf("tetromino [%s,%d]: dims expected {h:%d, w:%d}, got: {h:%d, w:%d}",
				t.name, r, p.height, p.width, len(p.cells), len(p.cells[0]))
			return
		}

		var hc [4]int
		for x := 0; x < p.width; x++ {
			for y := 0; y < p.height; y++ {
				if p.isFilled(y, x) {
					hc[x] = p.height - y
					break
				}
			}
		}
		for i, v := range p.heightCol {
			if v != hc[i] {
				test.Errorf("tetromino [%s,%d]: heightCol[%d] expected:%d, got:%d",
					t.name, r, i, v, hc[i])
				return
			}
		}

		n := 0
		for y, row := range p.cells {
			for x := range row {
				if p.isFilled(y, x) {
					n++
				}
			}
		}
		if n != 4 {
			test.Errorf("tetromino [%s,%d]: expected 4 cells, got: %d", t.name, r, n)
		}
	}
}

func TestTetrominoZ(test *testing.T) {
	tetrominoZ.TestTetromino(test)
}

func TestTetrominoI(test *testing.T) {
	tetrominoI.TestTetromino(test)
}
