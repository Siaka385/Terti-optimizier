package asfunc

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// CalculateBoardSize creates a square board with an optimal size for m tetriminos
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

// PlaceAllTetrimino attempts to place all tetriminos on the board
func PlaceAllTetrimino(board [][]string, tetri [][]string, index int) bool {
	if index == len(tetri) {
		return true
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if CanItBePlaced(board, tetri[index], i, j) {
				PlaceTetri(board, tetri[index], i, j, index)
				if PlaceAllTetrimino(board, tetri, index+1) {
					return true
				}
				RemoveTetri(board, tetri[index], i, j)
			}
		}
	}
	return false
}

// CanItBePlaced checks if a tetrimino can be placed at the given coordinates
func CanItBePlaced(board [][]string, tetri []string, x, y int) bool {
	for i, line := range tetri {
		for j, char := range line {
			if char == '#' {
				if x+i >= len(board) || y+j >= len(board[0]) || board[x+i][y+j] != "." {
					return false
				}
			}
		}
	}
	return true
}

// PlaceTetri places a tetrimino on the board
func PlaceTetri(board [][]string, tetri []string, x, y, letter int) {
	mySlice := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
		"Y", "Z",
	}

	for i, line := range tetri {
		for j, char := range line {
			if char == '#' {
				board[x+i][y+j] = mySlice[letter]
			}
		}
	}
}

// RemoveTetri removes a tetrimino from the board
func RemoveTetri(board [][]string, tetri []string, x, y int) {
	for i, line := range tetri {
		for j, char := range line {
			if char == '#' {
				board[x+i][y+j] = "."
			}
		}
	}
}

// PrintBoard prints the board to the console
func PrintBoard(board [][]string) {
	for _, row := range board {
		fmt.Println(strings.Join(row, ""))
	}
}

// Proo optimizes the board size and attempts to place all tetriminos
func Proo(tetri []string, m int) {
	start := time.Now()

	// Pre-split the tetriminos for optimization
	splitTetri := make([][]string, len(tetri))
	for i, t := range tetri {
		splitTetri[i] = strings.Split(t, "\n")
	}

	for {
		board := CalculateBoardSize(m)
		if PlaceAllTetrimino(board, splitTetri, 0) {
			fmt.Println("All Tetrimino pieces placed successfully:")
			PrintBoard(board)
			break
		}
		m++
	}

	stop := time.Since(start)
	fmt.Println("Time elapsed:", stop)
}

// Here’s an optimized version of your code, with some improvements for readability and efficiency:

//     Reduced Recalculations: The original code recalculates some values repeatedly. I’ve minimized this by calculating these values once where possible.
//     Removed Unnecessary Recursion: The Proo function recursively calls itself if all pieces aren’t placed. This can be made more efficient by iterating over the board size directly.
//     Reduced String Splitting: Splitting the tetrimino strings repeatedly was inefficient. I moved the splitting to the beginning to avoid redundant operations.

// 	Key Improvements:

//     Minimized Redundant Operations: The tetrimino strings are split once and stored, avoiding repetitive string operations.
//     Efficient Board Resizing: The resizing of the board is done iteratively, increasing the board size only if necessary.
//     Improved Time Logging: The time taken is logged after all operations are complete, providing a clearer picture of the total runtime.

// This should improve the efficiency and readability of your code while maintaining the same functionality.
