package player

import (
	"testing"
)

func (t *tetromino) testTetromino(test *testing.T) {
	for r, p := range t.pieces {

		if p.width != len(p.blocks[0]) || p.height != len(p.blocks) {
			test.Errorf("tetromino [%s,%d]: dims expected {h:%d, w:%d}, got: {h:%d, w:%d}",
				t.name, r, p.height, p.width, len(p.blocks), len(p.blocks[0]))
			return
		}

		var hc [4]int
		for x := 0; x < p.width; x++ {
			for y := 0; y < p.height; y++ {
				if p.blocks[y][x] == 1 {
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
		for y, row := range p.blocks {
			for x := range row {
				if p.blocks[y][x] == 1 {
					n++
				}
			}
		}
		if n != 4 {
			test.Errorf("tetromino [%s,%d]: expected 4 cells, got: %d", t.name, r, n)
		}
	}
}

func TestTetrominoI(test *testing.T) {
	tetrominoI.testTetromino(test)
}

func TestTetrominoJ(test *testing.T) {
	tetrominoJ.testTetromino(test)
}

func testTetrominoL(test *testing.T) {
	tetrominoL.testTetromino(test)
}

func TestTetrominoO(test *testing.T) {
	tetrominoO.testTetromino(test)
}

func TestTetrominoS(test *testing.T) {
	tetrominoS.testTetromino(test)
}

func TestTetrominoT(test *testing.T) {
	tetrominoT.testTetromino(test)
}

func TestTetrominoZ(test *testing.T) {
	tetrominoZ.testTetromino(test)
}