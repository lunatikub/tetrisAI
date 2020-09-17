package player

import (
	"testing"
)

func (f *field) testGetYPlay(test *testing.T, t *tetrimino, r int, x int, expectedY int) {
	p := &t.pieces[r]
	if y := f.getRow(p, x); y != expectedY {
		test.Errorf("tetrimino [%s,r:%d,x:%d]: expected:%d, got:%d",
			t.name, r, x, expectedY, y)
	}
}

func TestGetYPlay(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(18, []int{0, 1})
	f.setRow(19, []int{1, 1, 1})

	f.testGetYPlay(test, &tetriminoJ, 1, 0, 18)
	f.testGetYPlay(test, &tetriminoJ, 1, 1, 18)
	f.testGetYPlay(test, &tetriminoJ, 1, 2, 19)
	f.testGetYPlay(test, &tetriminoJ, 1, 3, 20)

	f.testGetYPlay(test, &tetriminoT, 2, 0, 18)
	f.testGetYPlay(test, &tetriminoT, 2, 1, 19)
	f.testGetYPlay(test, &tetriminoT, 2, 2, 20)
}

func TestGetYPlayWell(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(16, []int{1, 0, 1})
	f.setRow(17, []int{1, 0, 1})
	f.setRow(18, []int{1, 0, 1})
	f.setRow(19, []int{1, 0, 1})

	f.testGetYPlay(test, &tetriminoI, 1, 0, 16)
	f.testGetYPlay(test, &tetriminoI, 1, 1, 20)
	f.testGetYPlay(test, &tetriminoI, 1, 2, 16)
}

func Testput(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.put(&tetriminoJ.pieces[1], 0)
	f.put(&tetriminoO.pieces[0], 0)

	if !f.eqRow(15, []int{1, 1}) ||
		!f.eqRow(16, []int{1, 1}) ||
		!f.eqRow(17, []int{1, 0}) ||
		!f.eqRow(18, []int{1, 0}) ||
		!f.eqRow(19, []int{1, 1}) {
		test.Errorf("Testput row")
	}
	if f.col[0] != 5 || f.col[1] != 5 {
		test.Errorf("Testput col")
	}
}

func TestInvalidput(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	for i := 0; i < 5; i++ {
		if !f.put(&tetriminoI.pieces[1], 0) {
			test.Errorf("Expected put succeed")
		}
	}

	if f.put(&tetriminoI.pieces[1], 0) {
		test.Errorf("Expected put failed")
	}

	if !f.put(&tetriminoI.pieces[0], 6) {
		test.Errorf("Expected put succeed")
	}
	if f.put(&tetriminoI.pieces[0], 7) {
		test.Errorf("Expected put failed")
	}
}

func (f *field) fieldIsEmpty(test *testing.T) bool {
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

func TestClearRow(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.put(&tetriminoI.pieces[0], 0)
	f.put(&tetriminoI.pieces[0], 4)
	f.put(&tetriminoO.pieces[0], 8)
	f.put(&tetriminoI.pieces[0], 0)
	f.put(&tetriminoI.pieces[0], 4)
	if !f.fieldIsEmpty(test) {
		return
	}

	for x := 0; x < f.width; x++ {
		f.put(&tetriminoI.pieces[1], x)
	}
	if !f.fieldIsEmpty(test) {
		return
	}
}

func TestClearRowWithHole(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(18, []int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	f.setRow(19, []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 0})

	f.put(&tetriminoS.pieces[0], 0)

	if f.col[0] == 0 && f.col[1] == 2 && f.col[2] == 2 &&
		f.col[3] == 1 && f.col[4] == 1 && f.col[5] == 1 &&
		f.col[6] == 1 && f.col[7] == 1 && f.col[8] == 1 &&
		f.col[9] == 0 {
		return
	}
	test.Errorf("Test clear row with hole failed")
}
