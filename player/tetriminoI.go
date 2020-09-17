package player

var tetriminoI = tetrimino{
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
			[]int{1, 1, 1, 1},
			[]int{4},
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
			[]int{4},
			[]int{1, 1, 1, 1},
			4,
			1,
		},
	},
}
