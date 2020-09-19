package player

import (
	"testing"
)

func TestDeltaR(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(18, []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	f.setLine(19, []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

	expectedDr := 15
	if dr := f.deltaR(); dr != expectedDr {
		test.Errorf("DeltaR: expected:%d, got:%d", expectedDr, dr)
	}
}

func TestDeltaC(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(18, []int{0, 0, 1, 1, 1, 0, 1, 0, 1, 0})
	f.setLine(19, []int{0, 1, 0, 1, 0, 1, 0, 1, 0, 1})

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

	f.setLine(16, []int{1, 0, 0, 0, 1, 0, 0, 0, 0, 0})
	f.setLine(17, []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0})
	f.setLine(18, []int{0, 0, 1, 0, 1, 0, 0, 0, 0, 0})
	f.setLine(19, []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 0})

	expectedHo := 8
	if ho := f.holes(); ho != expectedHo {
		test.Errorf("Holes: expected:%d, got:%d", expectedHo, ho)
	}
}

func TestWells(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(16, []int{1})
	f.setLine(17, []int{1, 0, 1})
	f.setLine(18, []int{1, 0, 1, 0, 1, 0, 1})
	f.setLine(19, []int{1, 0, 1, 1, 1, 0, 1})

	expectedWe := 10
	if we := f.wells(); we != expectedWe {
		test.Errorf("Wells: expected:%d, got:%d", expectedWe, we)
	}
}

func TestErosion(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(18, []int{0, 0, 0, 1, 1, 1, 1, 1, 1, 1})
	f.setLine(19, []int{1, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	f.put(T, 2, 0)

	expected := 8
	if f.erosion != expected {
		test.Errorf("Erosion: expected:%d, got:%d", expected, f.erosion)
	}

	f.setLine(19, []int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	f.put(O, 0, 0)

	expected = 2
	if f.erosion != expected {
		test.Errorf("Erosion: expected:%d, got:%d", expected, f.erosion)
	}
}

func TestHLT(test *testing.T) {
	f := newField(fieldHeight, fieldWidth)

	f.setLine(19, []int{0, 0, 1, 1, 1, 1, 1, 1, 1, 1})
	f.put(T, 0, 1)

	expected := 1
	if f.hlt != expected {
		test.Errorf("Erosion: expected:%d, got:%d", expected, f.hlt)
	}
}

func TestEval(test *testing.T) {
	f1 := newField(fieldHeight, fieldWidth)
	f2 := newField(fieldHeight, fieldWidth)

	f1.setLine(19, []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0})
	f2.setLine(19, []int{1, 1, 1, 1, 1, 0, 0, 0, 0, 0})

	if f1.eval() > f2.eval() {
		test.Errorf("Eval: expected:%d less than %d", f1.eval(), f2.eval())
	}
}
