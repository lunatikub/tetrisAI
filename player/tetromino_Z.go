package player

var tetrominoZ = tetromino{
	"Z",
	[]piece{

		// +---+
		// |XX.|
		// |.XX|
		// +---+
		{
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
			},
			[]int{1, 0, 0},
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
			[]int{0, 1},
			3,
			2,
		},
	},
}
