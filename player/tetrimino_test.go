package player

import (
	"testing"
)

func (p *piece) testBlockRow(test *testing.T, name *string, r int) bool {
	return true
}

func (p *piece) testHC(test *testing.T, name *string, r int) bool {
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
			test.Errorf("tetrimino [%s,%d]: holes[%d] expected:%d, got:%d",
				*name, r, i, v, hc[i])
			return false
		}
	}
	return true
}

func (p *piece) testHoles(test *testing.T, name *string, r int) bool {
	var holes [4]int
	for x := 0; x < p.width; x++ {
		y := p.height - 1
		for {
			if y == 0 || p.blocks[y][x] == 1 {
				break
			}
			y--
		}
		holes[x] = p.height - y - 1
	}
	for i, v := range p.holes {
		if v != holes[i] {
			test.Errorf("tetrimino [%s,%d]: holes[%d] expected:%d, got:%d",
				*name, r, i, v, holes[i])
			return false
		}
	}
	return true
}

func (p *piece) testNrBlock(test *testing.T, name *string, r int) bool {
	n := 0
	for y, row := range p.blocks {
		for x := range row {
			if p.blocks[y][x] == 1 {
				n++
			}
		}
	}
	if n != 4 {
		test.Errorf("tetrimino [%s,%d]: expected 4 cells, got: %d", *name, r, n)
		return false
	}
	return true
}

func (p *piece) testDim(test *testing.T, name *string, r int) bool {
	if p.width != len(p.blocks[0]) || p.height != len(p.blocks) {
		test.Errorf("tetrimino [%s,%d]: dims expected {h:%d, w:%d}, got: {h:%d, w:%d}",
			*name, r, p.height, p.width, len(p.blocks), len(p.blocks[0]))
		return false
	}
	if p.width != len(p.holes) || p.width != len(p.heightCol) || p.height != len(p.blocksRow) {
		test.Errorf("tetrimino [%s,%d]: dims holes or heightCol or blocksRow unexpected", *name, r)
		return false
	}
	return true
}

func (t *tetrimino) testTetrimino(test *testing.T) {
	for r, p := range t.pieces {
		if !p.testDim(test, &t.name, r) {
			return
		}
		if !p.testHC(test, &t.name, r) {
			return
		}

		if !p.testHoles(test, &t.name, r) {
			return
		}

		if !p.testNrBlock(test, &t.name, r) {
			return
		}
	}
}

func TestTetriminoI(test *testing.T) {
	tetriminoI.testTetrimino(test)
}

func TestTetriminoJ(test *testing.T) {
	tetriminoJ.testTetrimino(test)
}

func testTetriminoL(test *testing.T) {
	tetriminoL.testTetrimino(test)
}

func TestTetriminoO(test *testing.T) {
	tetriminoO.testTetrimino(test)
}

func TestTetriminoS(test *testing.T) {
	tetriminoS.testTetrimino(test)
}

func TestTetriminoT(test *testing.T) {
	tetriminoT.testTetrimino(test)
}

func TestTetriminoZ(test *testing.T) {
	tetriminoZ.testTetrimino(test)
}
