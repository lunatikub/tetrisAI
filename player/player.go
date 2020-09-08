package player

// Player of tetris game
type Player struct {
	field field
}

// Dump debug dump function
func (p *Player) Dump() {
	p.field.dump()
}
