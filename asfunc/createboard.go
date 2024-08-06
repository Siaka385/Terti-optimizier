package asfunc

import (
	"fmt"
	"math"
	"strings"
)

func CalculateBoardSize(m int) [][]string {
	num := m * 4

	grid := int(math.Ceil(math.Sqrt(float64(num))))
	board := make([][]string, grid)

	for j := range board {
		board[j] = make([]string, grid)
		for i := range board[j] {
			board[j][i] = "."
		}
	}
	return board
}

func PlaceAllTertimino(board [][]string, tetri []string, index int) bool {
	if index == len(tetri) {
		return true
	}
	tetrimino := tetri[index]

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if CanItBePlaced(board, tetrimino, i, j) {
				Placetetri(board, tetrimino, i, j)
				if PlaceAllTertimino(board, tetri, index+1) {
					return true
				}
				Removetetri(board, tetrimino, i, j)
			}
		}
	}
	return false
}

func CanItBePlaced(board [][]string, tetri string, x, y int) bool {
	mytetri := strings.Split(tetri, "\n")
	for i, line := range mytetri {
		for j, char := range line {
			if char == '#' {
				// Check if the position is out of bounds or already occupied
				if x+i >= len(board) || y+j >= len(board[0]) || board[x+i][y+j] == "#" {
					return false
				}
			}
		}
	}
	return true
}

func Placetetri(board [][]string, tetri string, x, y int) {
	mytetri := strings.Split(tetri, "\n")
	for i, line := range mytetri {
		for j, char := range line {
			if char == '#' {
				board[x+i][y+j] = "#"
			}
		}
	}
}

func Removetetri(board [][]string, tetri string, x, y int) {
	mytetri := strings.Split(tetri, "\n")
	for i, line := range mytetri {
		for j, char := range line {
			if char == '#' {
				board[x+i][y+j] = "."
			}
		}
	}
}

func PrintBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}

func Proo(teti []string, m int) {
	bord := CalculateBoardSize(m)

	if PlaceAllTertimino(bord, teti, 0) {
		fmt.Println("All Tetrimino pieces placed successfully:")
		PrintBoard(bord)
	} else {
		Proo(teti, m+1)
	}
}
