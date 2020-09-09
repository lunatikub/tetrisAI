package player

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
