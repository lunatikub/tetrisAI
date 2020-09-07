package tetromino

// TetrominoZ inverted bias
var tetrominoZ = Definition{
	"Z",
	[]Piece{

		// +---+
		// |XX.|
		// |.XX|
		// +---+
		{
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
			},
			[]int{2, 2, 1},
			2,
			3,
		},

		// +--+
		// |.X|
		// |XX|
		// |X.|
		// +--+
		{
			[][]int{
				{0, 1},
				{1, 1},
				{1, 0},
			},
			[]int{2, 3},
			3,
			2,
		},
	},
}
