package player

import "fmt"

type field struct {
	blocks [][]int
	col    []int // height by column
	row    []int // number of blocks filled by row
	height int
	width  int
}

// Create a new field for the tetris game
func newField(h int, w int) field {
	var f field
	f.blocks = make([][]int, h)
	for i := range f.blocks {
		f.blocks[i] = make([]int, w)
	}
	f.col = make([]int, w)
	f.row = make([]int, h)
	f.width = w
	f.height = h
	return f
}

// Get the Y coordonate to set a piece
func (f *field) getY(p *piece, x int) int {
	y := 0
	for i := 0; i < p.width; i++ {
		if r := f.col[x+i] - p.holes[i]; r > y {
			y = r
		}
	}
	return f.height - y
}

func (f *field) completeRow(y int) {
	copy(f.blocks[1:], f.blocks[:y])
	for i := range f.col {
		f.col[i]--
	}
	copy(f.row[1:], f.row[:y])
}

func (f *field) push(p *piece, x int) bool {
	if x+p.width > f.width {
		return false
	}
	y := f.getY(p, x)
	if y-p.height < 0 {
		return false
	}

	for i, row := range p.blocks {
		k := i + y - p.height
		for j, v := range row {
			if v == 1 {
				f.blocks[k][j+x] = v
				f.row[k]++
				if f.height-k > f.col[j+x] {
					f.col[j+x] = f.height - k
				}
				if f.row[k] == f.width {
					f.completeRow(k)
				}
			}
		}
	}
	return true
}

func (f *field) dump() {
	fmt.Println()
	fmt.Println("    0 1 2 3 4 5 6 7 8 9")
	for y, row := range f.blocks {
		fmt.Printf("%2d| ", y)
		for _, v := range row {
			if v == 1 {
				fmt.Print("X ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Printf(" |%2d\n", f.row[y])
	}
	fmt.Print("   ")
	for _, v := range f.col {
		fmt.Printf("%2d", v)
	}
	fmt.Println()
}
