package main

import (
	"log"
	"os"

	tetris "github.com/lunatikub/tetrisBot/player"
)

const (
	nrNextTetriminoes = 5
	heightField       = 20
	widthField        = 10
)

func main() {
	file, err := os.OpenFile("logs.txt",
		os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	p := tetris.NewPlayer(heightField, widthField)
	emu := tetris.NewEmu(nrNextTetriminoes)
	//p.DisplayerInit()

	n := 0
	for {
		//p.DisplayClear()
		//p.DisplayTetrimino(emu.CurrT, emu.NextT)
		p.Play(emu.CurrT, emu.NextT)
		//p.Display()
		if n == 1200 {
			break
		}
		n++
		emu.Next()
		//p.Wait(0)
	}
}
