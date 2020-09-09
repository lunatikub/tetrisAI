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
