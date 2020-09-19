package player

import (
	"log"
	"time"

	gc "github.com/gbin/goncurses"
)

const (
	xShift = 2 // x shift
	yShift = 1 // y shift
)

func initColor() {
	if err := gc.StartColor(); err != nil {
		log.Fatal(err)
	}

	gc.InitPair(I, gc.C_CYAN, gc.C_BLACK)
	gc.InitPair(J, gc.C_WHITE, gc.C_BLACK)
	gc.InitPair(L, gc.C_BLUE, gc.C_BLACK)
	gc.InitPair(O, gc.C_YELLOW, gc.C_BLACK)
	gc.InitPair(S, gc.C_GREEN, gc.C_BLACK)
	gc.InitPair(T, gc.C_MAGENTA, gc.C_BLACK)
	gc.InitPair(Z, gc.C_RED, gc.C_BLACK)
}

// DisplayerInit initialize the displayer of the field
func (p *Player) DisplayerInit() {
	stdscr, err := gc.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer gc.End()

	if !gc.HasColors() {
		log.Fatal("Example requires a colour capable terminal")
	}

	initColor()
	gc.Echo(false)
	gc.CBreak(true)
	gc.Cursor(0)
	p.stdscr = stdscr
}

// DisplayClear Clear the screen
func (p *Player) DisplayClear() {
	p.stdscr.Clear()
}

// DisplayTetrimino Display the tetriminos to come
func (p *Player) DisplayTetrimino(currT int, nextT []int) {
	s := p.stdscr
	f := p.field
	xs := xShift + 14
	s.MovePrintf(2, f.width*2+xs, "Current tetrimino [%d]", currT)
	s.MovePrintf(3, f.width*2+xs, "Next tetriminos   [")
	for x, v := range nextT {
		s.MovePrintf(3, f.width*2+xs+19+x*2, "%d,", v)
	}
	s.MovePrintf(3, f.width*2+xs+2*len(nextT)+18, "]")
}

// Display display the board and all properties
func (p *Player) Display() {
	s := p.stdscr
	f := p.field

	xs := xShift
	ys := yShift

	for y := 0; y < f.height; y++ {
		s.MovePrintf(y+ys, xs, "%2d| ", f.height-y-1)
		s.MovePrintf(y+ys, xs+f.width*2+4, "| [%2d]", f.line[y])
	}
	xs += 4

	for x := 0; x < f.width; x++ {
		y := f.height + ys
		s.MovePrintf(y, x*2+xs, "%-2d", x)
		s.MovePrintf(y+2, x*2+xs, "%-2d", f.col[x])
	}
	s.MovePrint(f.height+ys+2, f.width*2+xs+2, "[height]")

	s.MovePrintf(6, f.width*2+xs+10, "lines:     [%10d]", p.nrLine)
	s.MovePrintf(7, f.width*2+xs+10, "score:     [%10d]", p.score)
	s.MovePrintf(8, f.width*2+xs+10, "tetriminos:[%10d]", p.nrTetrimino)
	s.MovePrintf(9, f.width*2+xs+10, "holds:     [%10d]", p.nrHold)

	for y, line := range f.blocks {
		for x, v := range line {
			if v != 0 {
				s.ColorOn(int16(v))
				s.MovePrintf(y+ys, x*2+xs, "%-2d", v)
				s.ColorOff(int16(v))
			} else {
				s.MovePrint(y+ys, x*2+xs, ". ")
			}
		}
	}
	s.Refresh()
}

// Wait TODO
func (p *Player) Wait(t int) {
	if t == 0 {
		p.stdscr.GetChar()
	} else {
		time.Sleep(time.Duration(t) * time.Microsecond)
	}
}
