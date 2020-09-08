package player

import (
	"testing"
)

const (
	fieldHeight = 20
	fieldWidth  = 10
)

func (f *field) testGetYPlay(test *testing.T, t *tetromino, r int, x int, expectedY int) {
	p := &t.pieces[r]
	if y := f.getY(p, x); y != expectedY {
		test.Errorf("tetromino [%s,r:%d,x:%d]: expected:%d, got:%d",
			t.name, r, x, expectedY, y)
	}
}

func TestGetYPlay(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	// .....
	// .X...
	// XXX..
	f.blocks[19][0] = 1
	f.blocks[19][1] = 1
	f.blocks[19][2] = 1
	f.blocks[18][1] = 1
	f.col[0] = 1
	f.col[1] = 2
	f.col[2] = 1

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
	f.col[7] = 4
	f.col[9] = 4

	// X
	// X
	// X
	// X
	f.testGetYPlay(test, &tetrominoI, 1, 7, 16)
	f.testGetYPlay(test, &tetrominoI, 1, 8, 20)
	f.testGetYPlay(test, &tetrominoI, 1, 9, 16)
}

func TestPush(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.push(&tetrominoJ.pieces[1], 0)
	f.push(&tetrominoO.pieces[0], 0)

	// 15| XX
	// 16| XX
	// 17| X.
	// 18| X.
	// 19| XX
	if f.blocks[19][0] == 1 && f.blocks[19][1] == 1 &&
		f.blocks[18][0] == 1 && f.blocks[18][1] == 0 &&
		f.blocks[17][0] == 1 && f.blocks[17][1] == 0 &&
		f.blocks[16][0] == 1 && f.blocks[16][1] == 1 &&
		f.blocks[15][0] == 1 && f.blocks[15][1] == 1 &&
		f.col[0] == 5 && f.col[1] == 5 {
		return
	}

	test.Errorf("TestPush")
}

func TestInvalidPush(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	for i := 0; i < 5; i++ {
		if !f.push(&tetrominoI.pieces[1], 0) {
			test.Errorf("Expected push succeed")
		}
	}

	if f.push(&tetrominoI.pieces[1], 0) {
		test.Errorf("Expected push failed")
	}

	if !f.push(&tetrominoI.pieces[0], 6) {
		test.Errorf("Expected push succeed")
	}
	if f.push(&tetrominoI.pieces[0], 7) {
		test.Errorf("Expected push failed")
	}
}

func (f *field) fieldIsReset(test *testing.T) bool {
	for x := 0; x < f.width; x++ {
		if f.col[x] != 0 {
			test.Errorf("col[%d] expected:0, got:%d", x, f.col[x])
			return false
		}
	}
	for y := 0; y < f.height; y++ {
		if f.row[y] != 0 {
			test.Errorf("row[%d] expected:0, got:%d", y, f.row[y])
			return false
		}
	}
	return true
}

func TestCompleteRow(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.push(&tetrominoI.pieces[0], 0)
	f.push(&tetrominoI.pieces[0], 4)
	f.push(&tetrominoO.pieces[0], 8)
	f.push(&tetrominoI.pieces[0], 0)
	f.push(&tetrominoI.pieces[0], 4)
	if !f.fieldIsReset(test) {
		return
	}

	for x := 0; x < f.width; x++ {
		f.push(&tetrominoI.pieces[1], x)
	}
	if !f.fieldIsReset(test) {
		return
	}
}
