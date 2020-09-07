package tetromino

// TetrominoI stick / bar / long
var TetrominoI = Definition{
	"I",
	[]Piece{

		// +----+
		// |XXXX|
		// +----+
		{
			[][]int{
				{1, 1, 1, 1},
			},
			[]int{1, 1, 1, 1},
			1,
			4,
		},

		// +-+
		// |X|
		// |X|
		// |X|
		// |X|
		// +-+
		{
			[][]int{
				{1},
				{1},
				{1},
				{1},
			},
			[]int{4},
			4,
			1,
		},
	},
}
