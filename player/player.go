package player

import (
	"fmt"
	"sort"

	gc "github.com/gbin/goncurses"
)

// Player of tetris game
type Player struct {
	field       *field
	hold        int
	nrHold      int
	score       int
	nrLine      int
	nrTetrimino int
	stdscr      *gc.Window
}

// NewPlayer Create a new player
func NewPlayer(h int, w int) *Player {
	p := new(Player)
	p.field = newField(h, w)
	return p
}

type combination struct {
	r     int // rotation
	x     int // x location
	f     *field
	score int
}

type byScore []combination

func (s byScore) Len() int           { return len(s) }
func (s byScore) Less(i, j int) bool { return s[i].score > s[j].score }
func (s byScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var n int

// Play Do the next move for the player
func (p *Player) Play(currT int, nextT []int) {
	t := getTetrimino(currT)
	var c []combination

	for r := range t.pieces {
		for x := 0; x < p.field.width; x++ {
			f := p.field.duplicate()
			if ret := f.put(currT, r, x); ret {
				c = append(c, combination{r, x, f, f.eval()})
			}
		}
	}

	sort.Sort(byScore(c))
	fmt.Println("play", currT, c[0].r, c[0].x)
	fmt.Println("n", n)
	n++
	p.field.put(currT, c[0].r, c[0].x)
	p.field.dump()
}
