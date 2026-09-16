package tetris

import (
	"math/rand"
	"testing"
)

func TestMoveDownLocksPieceAtBottom(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	model.Current = Piece{Type: I, X: 3, Y: boardHeight - 2}

	if model.dropOne() {
		t.Fatal("dropOne should lock a piece at the bottom")
	}

	occupied := 0
	for _, cell := range model.Board[boardHeight-1] {
		if cell != 0 {
			occupied++
		}
	}
	if occupied != 4 {
		t.Fatalf("occupied cells on bottom row = %d, want 4", occupied)
	}
}

func TestClearLinesUpdatesScoreAndLevel(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	for column := range boardWidth {
		model.Board[boardHeight-1][column] = I + 1
	}

	model.clearLines()

	if model.Lines != 1 {
		t.Fatalf("lines = %d, want 1", model.Lines)
	}
	if model.Score != 100 {
		t.Fatalf("score = %d, want 100", model.Score)
	}
	for _, cell := range model.Board[boardHeight-1] {
		if cell != 0 {
			t.Fatal("cleared row still contains a block")
		}
	}
}

func TestRotateChangesPieceWhenPositionIsValid(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	original := model.Current

	if !model.rotate() {
		t.Fatal("rotate should succeed in an empty board")
	}
	if model.Current.Rotation == original.Rotation {
		t.Fatal("rotation did not change")
	}
}

func TestHardDropLocksAndScoresDistance(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	model.Current = Piece{Type: O, X: 3, Y: 0}

	model.hardDrop()

	if model.Score == 0 {
		t.Fatal("hard drop should award points")
	}
	if model.Current.Y != 0 {
		t.Fatal("a new piece should spawn at the top after hard drop")
	}
}
