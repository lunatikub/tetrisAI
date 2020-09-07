package tetromino

import (
	"testing"
)

func (t *tetromino) TestTetromino(test *testing.T) {
	for _, p := range t.pieces {

		if p.width != len(p.cells[0]) || p.height != len(p.cells) {
			test.Errorf("tetromino %s: dims expected {h:%d, w:%d}, got: {h:%d, w:%d}",
				t.name, p.height, p.width, len(p.cells), len(p.cells[0]))
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
				test.Errorf("tetromino %s: heightCol[%d] expected:%d, got:%d",
					t.name, i, v, hc[i])
			}
		}
	}
}

func TestTetrominoZ(test *testing.T) {
	tetrominoZ.TestTetromino(test)
}
