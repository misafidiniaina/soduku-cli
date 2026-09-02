package logic

import (
	"testing"
	"time"
)

func TestScore(t *testing.T) {
	puzzle := [9][9]int{{1, 0, 0}, {0, 2, 0}}
	cells := puzzle
	cells[0][1] = 8
	cells[1][0] = -4

	if got := Score(cells, puzzle, 1, 3*time.Second); got != 47 {
		t.Fatalf("Score() = %d, want 47", got)
	}
}

func TestScoreNeverNegative(t *testing.T) {
	if got := Score([9][9]int{}, [9][9]int{}, 100, time.Hour); got != 0 {
		t.Fatalf("Score() = %d, want 0", got)
	}
}
