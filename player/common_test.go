package player

const (
	fieldHeight = 20
	fieldWidth  = 10
)

func (f *field) setLine(y int, line []int) {
	copy(f.blocks[y], line)
	n := 0
	for x, v := range line {
		h := f.height - y
		if v == 1 {
			n++
			if f.col[x] < h {
				f.col[x] = h
			}
		}
	}
	f.line[y] = n
}

func (f *field) eqLine(y int, line []int) bool {
	for x, v := range line {
		if f.blocks[y][x] != v {
			return false
		}
	}
	return true
}
