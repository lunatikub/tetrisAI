package player

// Erosion is process at push.
// Identifer: e

// Height of last tetrimino played is process at push.
// Identifier: l

// Process row transitions.
// Identifer: dr
func (f *field) deltaR() int {
	dr := 0
	for y := range f.blocks {
		for x := 0; x < f.width-1; x++ {
			// if sum=0: 2 empty blocks, if sum=2: 2 filled blocks
			if f.blocks[y][x]+f.blocks[y][x+1] == 1 {
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
			if f.blocks[y][x]+f.blocks[y+1][x] == 1 {
				dc++
			}
		}
	}
	return dc
}

// Process the holes.
// Identifier: L
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
// Identifier: W
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
