package firstfour

import (
	"fmt"
	"strings"
)

//Tictactoe comment
func Tictactoe() {

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[1][1] = "O"
	board[0][1] = "X"
	board[2][1] = "O"
	board[2][2] = "O"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
