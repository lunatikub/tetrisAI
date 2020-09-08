package player

const (
	fieldHeight = 20
	fieldWidth  = 10
)

type field struct {
	blocks [fieldHeight][fieldWidth]int
	height [fieldWidth]int
}
