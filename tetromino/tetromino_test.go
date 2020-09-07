package tetromino

import (
	"testing"
)

func (td *Definition) TestTetromino(test *testing.T) {
	for r, p := range td.Pieces {

		if p.Width != len(p.Blocks[0]) || p.Height != len(p.Blocks) {
			test.Errorf("tetromino [%s,%d]: dims expected {h:%d, w:%d}, got: {h:%d, w:%d}",
				td.Name, r, p.Height, p.Width, len(p.Blocks), len(p.Blocks[0]))
			return
		}

		var hc [4]int
		for x := 0; x < p.Width; x++ {
			for y := 0; y < p.Height; y++ {
				if p.Blocks[y][x] == 1 {
					hc[x] = p.Height - y
					break
				}
			}
		}
		for i, v := range p.HeightCol {
			if v != hc[i] {
				test.Errorf("tetromino [%s,%d]: heightCol[%d] expected:%d, got:%d",
					td.Name, r, i, v, hc[i])
				return
			}
		}

		n := 0
		for y, row := range p.Blocks {
			for x := range row {
				if p.Blocks[y][x] == 1 {
					n++
				}
			}
		}
		if n != 4 {
			test.Errorf("tetromino [%s,%d]: expected 4 cells, got: %d", td.Name, r, n)
		}
	}
}

func TestTetrominoI(test *testing.T) {
	TetrominoI.TestTetromino(test)
}

func TestTetrominoJ(test *testing.T) {
	TetrominoJ.TestTetromino(test)
}

func TestTetrominoL(test *testing.T) {
	TetrominoL.TestTetromino(test)
}

func TestTetrominoO(test *testing.T) {
	TetrominoO.TestTetromino(test)
}

func TestTetrominoS(test *testing.T) {
	TetrominoS.TestTetromino(test)
}

func TestTetrominoT(test *testing.T) {
	TetrominoT.TestTetromino(test)
}

func TestTetrominoZ(test *testing.T) {
	tetrominoZ.TestTetromino(test)
}
