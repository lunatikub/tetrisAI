package player

import (
	"testing"
)

func (f *field) testGetYPlay(test *testing.T, t *tetromino, r int, x int, expectedY int) {
	p := &t.pieces[r]
	if y := f.getYPlay(p, x); y != expectedY {
		test.Errorf("tetromino [%s,r:%d,x:%d]: expected:%d, got:%d",
			t.name, r, x, expectedY, y)
	}
}

func TestGetYPlay(test *testing.T) {
	var f field

	// .....
	// .X...
	// XXX..
	f.blocks[19][0] = 1
	f.blocks[19][1] = 1
	f.blocks[19][2] = 1
	f.blocks[18][1] = 1
	f.height[0] = 1
	f.height[1] = 2
	f.height[2] = 1

	// X.
	// X.
	// XX
	f.testGetYPlay(test, &tetrominoJ, 1, 0, 2)
	f.testGetYPlay(test, &tetrominoJ, 1, 1, 2)
	f.testGetYPlay(test, &tetrominoJ, 1, 2, 1)
	f.testGetYPlay(test, &tetrominoJ, 1, 3, 0)

	// XXX
	// .X.
	f.testGetYPlay(test, &tetrominoT, 2, 0, 2)
	f.testGetYPlay(test, &tetrominoT, 2, 1, 1)
	f.testGetYPlay(test, &tetrominoT, 2, 2, 0)

	// X.X
	// X.X
	// X.X
	// X.X
	f.blocks[19][7] = 1
	f.blocks[18][7] = 1
	f.blocks[17][7] = 1
	f.blocks[16][7] = 1
	f.blocks[19][9] = 1
	f.blocks[18][9] = 1
	f.blocks[17][9] = 1
	f.blocks[16][9] = 1
	f.height[7] = 4
	f.height[9] = 4

	// X
	// X
	// X
	// X
	f.testGetYPlay(test, &tetrominoI, 1, 7, 4)
	f.testGetYPlay(test, &tetrominoI, 1, 8, 0)
	f.testGetYPlay(test, &tetrominoI, 1, 9, 4)
}
