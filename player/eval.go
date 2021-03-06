package player

// Erosion is proceed at put.
// Identifer: e

// Height of last tetrimino played is proceed at put.
// Identifier: l

// Process line transitions.
// Identifer: dr
func (f *field) deltaR() int {
	dr := 0
	for y := range f.blocks {
		for x := 0; x < f.width-1; x++ {
			// if sum=0: 2 empty blocks, if sum=2: 2 filled blocks
			if f.blocks[y][x]+f.blocks[y][x+1] > 0 {
				dr++
			}
		}
	}
	return dr
}

// Process column transitions.
// Identifer: dc
func (f *field) deltaC() int {
	dc := 0
	for x := 0; x < f.width; x++ {
		for y := 0; y < f.height-1; y++ {
			// if sum=0: 2 empty blocks, if sum=2: 2 filled blocks
			if f.blocks[y][x]+f.blocks[y+1][x] > 0 {
				dc++
			}
		}
	}
	return dc
}

// Process the holes.
// Identifier: h
func (f *field) holes() int {
	ho := 0
	for x := 0; x < f.width; x++ {
		for y := f.height - f.col[x] + 1; y < f.height; y++ {
			if f.blocks[y][x] == 0 {
				ho++
			}
		}
	}
	return ho
}

// sumSeq(n) = 1 + 2 + 3 + ... + N
func sumSeq(n int) int {
	return (n * (n + 1)) / 2
}

// get the minimum height between both neighbors of the well
func (f *field) getMinHeightNeighbors(x int) int {
	l := 0 // height of the left neighbor
	r := 0 // height of the right neighbor
	if x != 0 {
		l = f.col[x-1]
	}
	if x != f.width-1 {
		r = f.col[x+1]
	}
	if l > r {
		return r
	}
	return l
}

// Process the whells.
// Identifier: w
func (f *field) wells() int {
	we := 0
	for x := 0; x < f.width; x++ {
		min := f.getMinHeightNeighbors(x)
		if min > f.col[x] {
			we += sumSeq(min - f.col[x])
		}
	}
	return we
}

const (
	lW  = -1263
	eW  = 660
	drW = -922
	dcW = -1977
	hW  = -1308
	wW  = -1049
)

// Dellacherie −l + e − dr − dc − 4h − w
func (f *field) eval() int {
	return lW*f.hlt +
		eW*f.erosion +
		drW*f.deltaR() +
		dcW*f.deltaC() +
		4*hW*f.holes() +
		wW*f.wells()
}
