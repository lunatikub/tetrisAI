package tetromino

// TetrominoS bias
var TetrominoS = Definition{
	"S",
	[]Piece{

		// +---+
		// |.XX|
		// |XX.|
		// +---+
		{
			[][]int{
				{0, 1, 1},
				{1, 1, 0},
			},
			[]int{1, 2, 2},
			2,
			3,
		},

		// +--+
		// |X.|
		// |XX|
		// |.X|
		// +--+
		{
			[][]int{
				{1, 0},
				{1, 1},
				{0, 1},
			},
			[]int{3, 2},
			3,
			2,
		},
	},
}
