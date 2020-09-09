package player

import (
	"testing"
)

func TestDeltaR(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	// ..XXX.X.X.
	// .X.X.X.X.X
	copy(f.blocks[18], []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	copy(f.blocks[19], []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

	expectedDr := 15
	if dr := f.deltaR(); dr != expectedDr {
		test.Errorf("DeltaR: expected:%d, got:%d", expectedDr, dr)
	}
}

func TestDeltaC(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	// ..XXX.X.X.
	// .X.X.X.X.X
	copy(f.blocks[18], []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	copy(f.blocks[19], []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

	expectedDc := 13
	if dc := f.deltaC(); dc != expectedDc {
		test.Errorf("DeltaC: expected:%d, got:%d", expectedDc, dc)
	}
}

func TestHoles(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	// X...X
	// .X...
	// ..X.X
	// ...X.
	copy(f.blocks[16], []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0})
	copy(f.blocks[17], []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0})
	copy(f.blocks[18], []int{0, 0, 1, 0, 1, 0, 0, 0, 0, 0})
	copy(f.blocks[19], []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 0})

	expectedHo := 8
	if ho := f.holes(); ho != expectedHo {
		test.Errorf("Holes: expected:%d, got:%d", expectedHo, ho)
	}
}
