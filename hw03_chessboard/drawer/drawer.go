package chessboarddrawer

import (
	"fmt"
	"strings"
)

func DrawChessBoard(size int) {
	var boardString strings.Builder

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			switch {
			case i%2 == 0 && j%2 == 0, i%2 != 0 && j%2 != 0:
				boardString.WriteString("#")
			case i%2 == 0 && j%2 != 0, i%2 != 0 && j%2 == 0:
				boardString.WriteString(" ")
			}
		}
		boardString.WriteString("\n")
	}

	fmt.Print(boardString.String())
}
