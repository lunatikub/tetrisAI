package player

import (
	"testing"
)

func (f *field) testGetYPlay(test *testing.T, t *tetrimino, r int, x int, expectedY int) {
	p := &t.pieces[r]
	if y := f.getLine(p, x); y != expectedY {
		test.Errorf("tetrimino [%s,r:%d,x:%d]: expected:%d, got:%d",
			t.name, r, x, expectedY, y)
	}
}

func TestGetYPlay(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(18, []int{0, 1})
	f.setLine(19, []int{1, 1, 1})

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

	f.setLine(16, []int{1, 0, 1})
	f.setLine(17, []int{1, 0, 1})
	f.setLine(18, []int{1, 0, 1})
	f.setLine(19, []int{1, 0, 1})

	f.testGetYPlay(test, &tetriminoI, 1, 0, 16)
	f.testGetYPlay(test, &tetriminoI, 1, 1, 20)
	f.testGetYPlay(test, &tetriminoI, 1, 2, 16)
}

func TestPut(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.put(J, 1, 0)
	f.put(O, 0, 0)

	if !f.eqLine(15, []int{1, 1}) ||
		!f.eqLine(16, []int{1, 1}) ||
		!f.eqLine(17, []int{1, 0}) ||
		!f.eqLine(18, []int{1, 0}) ||
		!f.eqLine(19, []int{1, 1}) {
		test.Errorf("TestPut line")
	}
	if f.col[0] != 5 || f.col[1] != 5 {
		test.Errorf("TestPut col")
	}
}

func TestInvalidput(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	for i := 0; i < 5; i++ {
		if !f.put(I, 1, 0) {
			test.Errorf("Expected put succeed")
		}
	}

	if f.put(I, 1, 0) {
		test.Errorf("Expected put failed")
	}

	if !f.put(I, 0, 6) {
		test.Errorf("Expected put succeed")
	}
	if f.put(I, 0, 7) {
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
		if f.line[y] != 0 {
			test.Errorf("line[%d] expected:0, got:%d", y, f.line[y])
			return false
		}
	}
	return true
}

func TestClearLine(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.put(I, 0, 0)
	f.put(I, 0, 4)
	f.put(O, 0, 8)
	f.put(I, 0, 0)
	f.put(I, 0, 4)
	if !f.fieldIsEmpty(test) {
		return
	}

	for x := 0; x < f.width; x++ {
		f.put(I, 1, x)
	}
	if !f.fieldIsEmpty(test) {
		return
	}
}

func TestClearLineWithHole(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(18, []int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	f.setLine(19, []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 0})

	f.put(S, 0, 0)

	if f.col[0] == 0 && f.col[1] == 2 && f.col[2] == 2 &&
		f.col[3] == 1 && f.col[4] == 1 && f.col[5] == 1 &&
		f.col[6] == 1 && f.col[7] == 1 && f.col[8] == 1 &&
		f.col[9] == 0 {
		return
	}
	test.Errorf("Test clear line with hole failed")
}

func TestBug513(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(14, []int{0, 0, 0, 5, 0, 0, 0, 0, 0, 0})
	f.setLine(15, []int{0, 0, 0, 5, 5, 2, 2, 0, 0, 0})
	f.setLine(16, []int{0, 7, 3, 3, 5, 7, 2, 0, 5, 5})
	f.setLine(17, []int{7, 7, 3, 3, 7, 7, 2, 5, 5, 0})
	f.setLine(18, []int{0, 3, 5, 4, 4, 2, 7, 3, 3, 2})
	f.setLine(19, []int{0, 3, 3, 3, 2, 7, 7, 3, 2, 5})

	f.dump()
	f.put(5, 1, 3)
	f.dump()
}
