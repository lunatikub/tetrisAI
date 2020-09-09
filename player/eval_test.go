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

	if dr := f.deltaR(); dr != 15 {
		test.Errorf("DeltaR: expected:15, got:%d", dr)
	}
}

func TestDeltaC(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	// ..XXX.X.X.
	// .X.X.X.X.X
	copy(f.blocks[18], []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	copy(f.blocks[19], []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

	if dc := f.deltaC(); dc != 13 {
		test.Errorf("DeltaC: expected:13, got:%d", dc)
	}
}
