package tetris

import (
	"github.com/charmbracelet/lipgloss"
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

func TestBagContainsEveryPiece(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.bag = nil
	for batch := 0; batch < 5; batch++ {
		seen := map[PieceType]bool{}
		for i := 0; i < 7; i++ {
			seen[m.newPiece().Type] = true
		}
		if len(seen) != 7 {
			t.Fatal("bag omitted a tetromino")
		}
	}
}

func TestPausedRotationAndWallKick(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Paused = true
	original := m.Current
	if m.rotate() || m.Current != original {
		t.Fatal("rotated while paused")
	}
	m.Paused = false
	m.Current = Piece{Type: I, Rotation: 1, X: -2}
	if !m.valid(m.Current) || !m.rotate() || !m.valid(m.Current) {
		t.Fatal("wall kick failed")
	}
}

func TestLandingPreviewMatchesHardDrop(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Current = Piece{Type: T, X: 3}
	ghost := m.landingPiece()
	m.hardDrop()
	for _, p := range m.cells(ghost) {
		if m.Board[p.Y][p.X] != T+1 {
			t.Fatal("preview does not match locked piece")
		}
	}
}

func TestLevelBoundaryUsesPreviousLevelForScore(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Lines = 9
	for x := range boardWidth {
		m.Board[boardHeight-1][x] = I + 1
	}
	m.clearLines()
	if m.Score != 100 || m.Level != 2 {
		t.Fatal("incorrect level-boundary score")
	}
}

func TestViewFitsStandardTerminal(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	for _, state := range []string{"playing", "paused", "over"} {
		m.Paused = state == "paused"
		m.GameOver = state == "over"
		view := m.View()
		if lipgloss.Height(view) > 24 || lipgloss.Width(view) > 80 {
			t.Fatalf("%s view exceeds 80x24: %dx%d", state, lipgloss.Width(view), lipgloss.Height(view))
		}
	}
}
