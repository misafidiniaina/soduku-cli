package gen

import "testing"

func TestPuzzleGen(t *testing.T) {
	puzzle, solution := PuzzleGen()
	if !IsSodukuValid(solution) {
		t.Fatal("generated solution is not a valid sudoku")
	}
	if !IsSodukuValid(puzzle) {
		t.Fatal("generated puzzle is not a valid partial sudoku")
	}

	empty := 0
	for row := range 9 {
		for column := range 9 {
			if puzzle[row][column] == 0 {
				empty++
				continue
			}
			if puzzle[row][column] != solution[row][column] {
				t.Fatalf("puzzle value at (%d,%d) differs from solution", row, column)
			}
		}
	}
	if empty != 45 {
		t.Fatalf("generated puzzle has %d empty cells, want 45", empty)
	}
}
