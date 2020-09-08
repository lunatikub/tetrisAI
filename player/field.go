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

func (f *field) getYPlay(p *piece, x int) int {
	y := 0
	for i := 0; i < p.width; i++ {
		if r := f.height[x+i] - p.holes[i]; r > y {
			y = r
		}
	}
	return y
}

func (f *field) dump() {
	for _, row := range f.blocks {
		for _, v := range row {
			if v == 1 {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}
