package player

import (
	"testing"
)

func (f *field) setRow(y int, row []int) {
	for x := range row {
		f.blocks[y][x] = row[x]
	}
}

func TestDeltaR(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	// ..XXX.X.X.
	// .X.X.X.X.X
	f.setRow(18, []int{0011101010})
	f.setRow(19, []int{0101010101})

	if dr := f.deltaR(); dr != 15 {
		test.Errorf("Expected DR:15, got DR:%d", dr)
	}
}
