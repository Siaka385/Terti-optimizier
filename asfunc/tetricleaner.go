package asfunc

import (
	"strings"
)

// Function to clean rows and columns with no '#'
func Clean(m []string) []string {
	var cleaned []string

	for _, line := range m {
		cleaned = append(cleaned, cleans(line))
	}
	return cleaned
}

// Function to clean a single tetrimino
func cleans(m string) string {
	mytetr := strings.Split(m, "\n")
	var mytetti [][]string
	for _, line := range mytetr {
		mytetti = append(mytetti, strings.Split(line, ""))
	}

	// Remove empty rows
	var cleanedRows [][]string
	// columns to print
	var columnSet []int
	for _, row := range mytetti {
		hasHash := false
		for j, char := range row {
			if char == "#" {
				hasHash = true
				columnSet = append(columnSet, j)
				break
			}
		}
		if hasHash {
			cleanedRows = append(cleanedRows, row)
		}
	}
	// Transpose to remove empty columns
	transposit := transpose(cleanedRows)

	// Remove empty columns
	var cleanedCols [][]string
	for _, col := range transposit {
		hasHash := false
		for _, char := range col {
			if char == "#" {
				hasHash = true
				break
			}
		}
		if hasHash {
			cleanedCols = append(cleanedCols, col)
		}
	}

	// Transpose back to original orientation
	cleaned := transpose(cleanedCols)

	// // Convert cleaned 2D slice back to string
	var result []string
	for _, row := range cleaned {
		result = append(result, strings.Join(row, ""))
	}

	return strings.Join(result, "\n")
}

// Function to transpose a 2D slice of strings
func transpose(m [][]string) [][]string {
	numRow := len(m)
	numCol := len(m[0])

	transp := make([][]string, numCol) // Note the change: numCol instead of numRow
	for i := range transp {
		transp[i] = make([]string, numRow)
	}
	for i := 0; i < numRow; i++ {
		for j := 0; j < numCol; j++ {
			transp[j][i] = m[i][j]
		}
	}
	return transp
}
