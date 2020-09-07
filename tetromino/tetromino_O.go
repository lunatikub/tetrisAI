package tetromino

// TetrominoO square / block
var TetrominoO = Definition{
	"O",
	[]Piece{

		// +--+
		// |XX|
		// |XX|
		// +--+
		{
			[][]int{
				{1, 1},
				{1, 1},
			},
			[]int{2, 2},
			2,
			2,
		},
	},
}
