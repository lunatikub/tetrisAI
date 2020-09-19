package player

import (
	"fmt"
)

type field struct {
	blocks  [][]int
	col     []int // height by column
	line    []int // number of blocks filled by line
	height  int
	width   int
	hlt     int
	erosion int
}

// Create a new field for the tetris game
func newField(h int, w int) *field {
	f := new(field)
	f.blocks = make([][]int, h)
	for i := range f.blocks {
		f.blocks[i] = make([]int, w)
	}
	f.col = make([]int, w)
	f.line = make([]int, h)
	f.width = w
	f.height = h
	return f
}

func (f *field) duplicate() *field {
	newF := newField(f.height, f.width)
	for i := range f.blocks {
		copy(newF.blocks[i], f.blocks[i])
	}
	copy(newF.col, f.col)
	copy(newF.line, f.line)
	return newF
}

// Get the line how to put a piece
func (f *field) getLine(p *piece, x int) int {
	y := 0
	for i := 0; i < p.width; i++ {
		if r := f.col[x+i] - p.holes[i]; r > y {
			y = r
		}
	}
	return f.height - y
}

func (f *field) setCol() {
	for x := 0; x < f.width; x++ {
		y := 0
		for {
			if y == f.height || f.blocks[y][x] > 0 {
				break
			}
			y++
		}
		f.col[x] = f.height - y
	}
}

func (f *field) updateLine(y int) bool {
	f.line[y]++
	// clear the line when it is full
	if f.line[y] == f.width {
		copy(f.blocks[1:], f.blocks[:y])
		copy(f.line[1:], f.line[:y])
		f.setCol()
		return true
	}
	return false
}

func (f *field) updateCol(x int, h int) {
	if f.height-h > f.col[x] {
		f.col[x] = f.height - h
	}
}

func (f *field) put(t int, r int, x int) bool {
	p := getPiece(t, r)
	if x+p.width > f.width {
		return false
	}
	y := f.getLine(p, x)
	if y-p.height < 0 {
		return false
	}

	nrCompletedLine := 0
	nrCompletedCell := 0
	for i, line := range p.blocks {
		k := i + y - p.height
		for j, v := range line {
			if v != 0 {
				f.blocks[k][j+x] = t
				f.updateCol(j+x, k)
				if f.updateLine(k) {
					nrCompletedLine++
					nrCompletedCell += p.blockLine[i]
				}
			}
		}
	}

	f.hlt = f.height - y
	f.erosion = nrCompletedLine * nrCompletedCell
	return true
}

func (f *field) dump() {
	for _, line := range f.blocks {
		for _, v := range line {
			if v != 0 {
				fmt.Printf("%d", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	for _, v := range f.col {
		fmt.Print(v, " ")
	}
	fmt.Println()
	for _, v := range f.line {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
