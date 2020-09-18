package player

import "fmt"

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
func newField(h int, w int) field {
	var f field
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
			if y == f.height || f.blocks[y][x] == 1 {
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

func (f *field) put(p *piece, x int) bool {
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
			if v == 1 {
				f.blocks[k][j+x] = v
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
	fmt.Println()
	fmt.Println("    0 1 2 3 4 5 6 7 8 9")
	for y, line := range f.blocks {
		fmt.Printf("%2d| ", y)
		for _, v := range line {
			if v == 1 {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Printf(" |%2d\n", f.line[y])
	}
	fmt.Print("   ")
	for _, v := range f.col {
		fmt.Printf("%2d", v)
	}
	fmt.Println()
}
