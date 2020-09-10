package player

const (
	fieldHeight = 20
	fieldWidth  = 10
)

func (f *field) setRow(y int, row []int) {
	copy(f.blocks[y], row)
	n := 0
	for x, v := range row {
		h := f.height - y
		if v == 1 {
			n++
			if f.col[x] < h {
				f.col[x] = h
			}
		}
	}
	f.row[y] = n
}

func (f *field) eqRow(y int, row []int) bool {
	for x, v := range row {
		if f.blocks[y][x] != v {
			return false
		}
	}
	return true
}
