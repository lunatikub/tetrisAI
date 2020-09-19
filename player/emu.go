package player

import (
	"fmt"
	"math/rand"
	"time"
)

// Emu emulator of a tetris game
type Emu struct {
	CurrT int   // current tetrimino
	NextT []int // next tetriminoes
}

func getRdmTetrimino() int {
	return rand.Intn(nrTetrimino-1) + 1
}

// NewEmu Create an emulator of a tetris game
func NewEmu(n int) *Emu {
	emu := new(Emu)
	rand.Seed(time.Now().UnixNano())
	emu.NextT = make([]int, n)
	emu.CurrT = getRdmTetrimino()
	for x := range emu.NextT {
		emu.NextT[x] = getRdmTetrimino()
	}
	return emu
}

// Next Emulate next round
func (e *Emu) Next() {
	e.CurrT = e.NextT[0]
	copy(e.NextT[0:], e.NextT[1:])
	e.NextT[len(e.NextT)-1] = getRdmTetrimino()
}

// Dump the current tetrimino and the next tetriminoes
func (e *Emu) Dump() {
	fmt.Println("Current tetrimino:", e.CurrT)
	fmt.Print("Next tetriminoes: [")
	for _, t := range e.NextT {
		fmt.Print(t, ",")
	}
	fmt.Println("]")
}
