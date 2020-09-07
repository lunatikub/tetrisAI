package main

import (
	"fmt"

	field "github.com/lunatikub/tetrisAI/field"
	tetromino "github.com/lunatikub/tetrisAI/tetromino"
)

func main() {
	var f field.Field

	f.Blocks[19][0] = 1
	n := f.GetHeight(&tetromino.TetrominoO.Pieces[0], 0)
	fmt.Println(n)
}
