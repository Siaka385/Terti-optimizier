package asfunc

import (
	"math"
	"testing"
	"time"
)

func TestValid4x4GridWithCorrectConnections(t *testing.T) {
	input := "####\n....\n....\n...."
	result := IsconnectedAtJoints(input)
	if !result {
		t.Errorf("Expected true, got %v", result)
	}
}

func TestGridWithLessThan4Lines(t *testing.T) {
	input := "###\n.#.\n..."
	result := IsconnectedAtJoints(input)
	if result {
		t.Errorf("Expected false, got %v", result)
	}
}

func TestJanitorMultipleLinesWithHashes(t *testing.T) {
	input := "##\n#.\n.#"
	expected := "##\n#.\n.#"
	result := janitor(input)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestJanitorSingleLine(t *testing.T) {
	input := "##..\n##..\n"
	expected := "##\n##"
	result := janitor(input)
	if result != expected {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestCalculateBoardSizeWithValidInput(t *testing.T) {
	m := 5
	expectedSize := int(math.Ceil(math.Sqrt(float64(m * 4))))
	board := CalculateBoardSize(m)

	if len(board) != expectedSize {
		t.Errorf("Expected board size %d, but got %d", expectedSize, len(board))
	}

	for _, row := range board {
		if len(row) != expectedSize {
			t.Errorf("Expected row size %d, but got %d", expectedSize, len(row))
		}
	}
}

func TestCalculateBoardSizeWithZeroInput(t *testing.T) {
	m := 0
	expectedSize := 0
	board := CalculateBoardSize(m)

	if len(board) != expectedSize {
		t.Errorf("Expected board size %d, but got %d", expectedSize, len(board))
	}

	for _, row := range board {
		if len(row) != expectedSize {
			t.Errorf("Expected row size %d, but got %d", expectedSize, len(row))
		}
	}
}

func TestProoExecutesWithoutErrors(t *testing.T) {
	tetri := []string{
		"####\n....\n....\n....",
		"##..\n##..\n....\n....",
	}
	m := 4

	start := time.Now()
	Proo(tetri, m)
	stop := time.Since(start)

	if stop <= 0 {
		t.Errorf("Expected positive time elapsed, got %v", stop)
	}
}
