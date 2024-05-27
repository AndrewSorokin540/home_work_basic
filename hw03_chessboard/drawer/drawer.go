package chessboarddrawer

import "fmt"

func DrawChessBoard(size int) {
	var boardString string

	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			switch {
			case i%2 == 0 && j%2 == 0, i%2 != 0 && j%2 != 0:
				boardString += string('#')
			case i%2 == 0 && j%2 != 0, i%2 != 0 && j%2 == 0:
				boardString += string('_')
			}
		}
		boardString += string('\n')
	}

	fmt.Print(boardString)
}
