package player

import "fmt"

// Player of tetris game
type Player struct {
	field field
}

// DumpField Dump the field
func (p *Player) DumpField() {
	for _, row := range p.field.blocks {
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
