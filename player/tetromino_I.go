package player

var tetrominoI = tetromino{
	"I",
	[]piece{

		// +----+
		// |XXXX|
		// +----+
		{
			[][]int{
				{1, 1, 1, 1},
			},
			[]int{0, 0, 0, 0},
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
			[]int{0},
			4,
			1,
		},
	},
}
