package player

import (
	"testing"
)

func TestDeltaR(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(18, []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	f.setRow(19, []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

	expectedDr := 15
	if dr := f.deltaR(); dr != expectedDr {
		test.Errorf("DeltaR: expected:%d, got:%d", expectedDr, dr)
	}
}

func TestDeltaC(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(18, []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	f.setRow(19, []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

	expectedDc := 13
	if dc := f.deltaC(); dc != expectedDc {
		const (
			fieldHeight = 20
			fieldWidth  = 10
		)
		test.Errorf("DeltaC: expected:%d, got:%d", expectedDc, dc)
	}
}

func TestHoles(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(16, []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0})
	f.setRow(17, []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0})
	f.setRow(18, []int{0, 0, 1, 0, 1, 0, 0, 0, 0, 0})
	f.setRow(19, []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 0})

	expectedHo := 8
	if ho := f.holes(); ho != expectedHo {
		test.Errorf("Holes: expected:%d, got:%d", expectedHo, ho)
	}
}

func TestWells(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setRow(16, []int{1})
	f.setRow(17, []int{1, 0, 1})
	f.setRow(18, []int{1, 0, 1, 0, 1})
	f.setRow(19, []int{1, 0, 1, 0, 1, 0, 1})

	expectedWe := 10
	if we := f.wells(); we != expectedWe {
		test.Errorf("Wells: expected:%d, got:%d", expectedWe, we)
	}
}
