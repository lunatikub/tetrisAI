package player

import (
	"testing"
)

func (f *field) testGetYPlay(test *testing.T, t *tetromino, r int, x int, expectedY int) {
	p := &t.pieces[r]
	if y := f.getY(p, x); y != expectedY {
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
	f.testGetYPlay(test, &tetrominoJ, 1, 0, 18)
	f.testGetYPlay(test, &tetrominoJ, 1, 1, 18)
	f.testGetYPlay(test, &tetrominoJ, 1, 2, 19)
	f.testGetYPlay(test, &tetrominoJ, 1, 3, 20)

	// XXX
	// .X.
	f.testGetYPlay(test, &tetrominoT, 2, 0, 18)
	f.testGetYPlay(test, &tetrominoT, 2, 1, 19)
	f.testGetYPlay(test, &tetrominoT, 2, 2, 20)

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
	f.testGetYPlay(test, &tetrominoI, 1, 7, 16)
	f.testGetYPlay(test, &tetrominoI, 1, 8, 20)
	f.testGetYPlay(test, &tetrominoI, 1, 9, 16)
}

func TestPush(test *testing.T) {
	var f field

	f.push(&tetrominoJ.pieces[1], 0)
	f.push(&tetrominoO.pieces[0], 0)

	// 15| XX
	// 16| XX
	// 17| X.
	// 18| X.
	// 19| XX
	//     01
	if f.blocks[19][0] == 1 && f.blocks[19][1] == 1 &&
		f.blocks[18][0] == 1 && f.blocks[18][1] == 0 &&
		f.blocks[17][0] == 1 && f.blocks[17][1] == 0 &&
		f.blocks[16][0] == 1 && f.blocks[16][1] == 1 &&
		f.blocks[15][0] == 1 && f.blocks[15][1] == 1 {
		return
	}
	test.Errorf("TestPush")
}
