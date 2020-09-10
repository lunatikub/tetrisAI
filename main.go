package main

import (
	tetris "github.com/lunatikub/tetrisAI/player"
)

const (
	nrNextTetrominoes = 5
	heightField       = 20
	widthField        = 10
)

func main() {
	p := tetris.NewPlayer(heightField, widthField)
	emu := tetris.NewEmu(nrNextTetrominoes)

	n := 0

	for {
		emu.Dump()
		p.Move(emu.CurrT, emu.NextT)
		emu.Next()
		if n == 20 {
			break
		}
		n++
	}

}
