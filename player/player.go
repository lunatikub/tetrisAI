package player

import "fmt"

// Player of tetris game
type Player struct {
	field field
	hold  int
}

// NewPlayer Create a new player
func NewPlayer(h int, w int) *Player {
	p := new(Player)
	p.field = newField(h, w)
	return p
}

// Move Do the next move for the player
func (p *Player) Move(currT int, nextT []int) {
	t := getTetromino(currT)
	fmt.Println("move", t.name)
	for _, piece := range t.pieces {
		for x := 0; x < p.field.width-piece.width+1; x++ {
			p.field.push(&piece, x)
			p.field.dump()
		}
	}
}

// Dump debug dump function
func (p *Player) Dump() {
	p.field.dump()
}
