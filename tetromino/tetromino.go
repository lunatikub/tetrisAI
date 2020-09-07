package tetromino

const (
	maxRotation = 4
	nrRow       = 4
	nrCol       = 4
)

// Tetromino rotation
type tRot [nrRow][nrCol]int

// Tetromino definition
type tDef struct {
	rot tRot        // Tetromino rotation
	hc  [nrCol]uint // Height by column
	h   uint        // Height of the tetromino
	w   uint        // Width of the tetromino
}

// Tetromino is a game piece
type Tetromino struct {
	nrRotation uint              // number of rotation
	defs       [maxRotation]tDef // tetromino rotations
}
