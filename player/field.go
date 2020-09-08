package player

import "fmt"

const (
	fieldHeight = 20
	fieldWidth  = 10
)

type field struct {
	blocks [fieldHeight][fieldWidth]int
	height [fieldWidth]int
}

func (f *field) getY(p *piece, x int) int {
	y := 0
	for i := 0; i < p.width; i++ {
		if r := f.height[x+i] - p.holes[i]; r > y {
			y = r
		}
	}
	return fieldHeight - y
}

func (f *field) push(p *piece, x int) bool {
	y := f.getY(p, x)

	if x+p.width > fieldWidth || y-p.height < 0 {
		return false
	}

	for i, row := range p.blocks {
		for j, v := range row {
			f.blocks[i+y-p.height][j+x] = v
		}
	}
	f.height[x] += p.height

	return true
}

func (f *field) dump() {
	for y, row := range f.blocks {
		fmt.Printf("%2d| ", y)
		for _, v := range row {
			if v == 1 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
	fmt.Println("    0123456789")
}
