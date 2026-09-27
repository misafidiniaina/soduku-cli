package game2048

import (
	"math/rand"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
	"github.com/muesli/termenv"
)

func press(m Model, key string) Model {
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)})
	return updated.(Model)
}

func countTiles(board Board) int {
	count := 0
	for _, row := range board {
		for _, value := range row {
			if value != 0 {
				count++
			}
		}
	}
	return count
}

func TestStartAndRestart(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	if countTiles(m.Board) != 2 || m.Board.highest() > 4 {
		t.Fatal("new game must start with two 2-or-4 tiles")
	}
	m.Score, m.BestScore, m.Moves, m.Width = 100, 200, 10, 120
	m.Paused, m.GameOver, m.Won, m.KeepPlaying = true, true, true, true
	m = press(m, "r")
	if countTiles(m.Board) != 2 || m.Score != 0 || m.Moves != 0 || m.Paused || m.GameOver || m.Won || m.KeepPlaying {
		t.Fatal("restart did not reset the game")
	}
	if m.BestScore != 200 || m.Width != 120 {
		t.Fatal("restart should retain session best and terminal dimensions")
	}
}

func TestOnlyChangedMovesSpawnAndScore(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Board = Board{{2, 4, 8, 16}}
	before := m.Board
	if m.move(left) || m.Board != before || m.Score != 0 || m.Moves != 0 {
		t.Fatal("blocked move changed the game")
	}
	m.Board = Board{{2, 2}}
	if !m.move(left) || m.Board[0][0] != 4 || countTiles(m.Board) != 2 || m.Score != 4 || m.BestScore != 4 || m.Moves != 1 {
		t.Fatal("merge should score once and spawn exactly one new tile")
	}
}

func TestWinAndContinue(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Board = Board{{1024, 1024}}
	m = press(m, "a")
	if !m.Won || m.Score != 2048 || m.Board.highest() != 2048 {
		t.Fatal("winning merge was not recognized")
	}
	before := m.Board
	m = press(m, "d")
	if m.Board != before {
		t.Fatal("win screen must wait for confirmation")
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = press(updated.(Model), "d")
	if !m.KeepPlaying || m.Board == before || m.Moves != 2 {
		t.Fatal("could not continue beyond 2048")
	}
}

func TestFinalSpawnCanEndGame(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Board = Board{{2, 2, 8, 16}, {8, 16, 8, 16}, {16, 8, 16, 8}, {8, 16, 8, 16}}
	if !m.move(left) || !m.GameOver || m.Score != 4 {
		t.Fatal("game should end when the final spawn leaves no legal moves")
	}
	before := m.Board
	m = press(m, "d")
	if m.Board != before {
		t.Fatal("game-over board accepted a move")
	}
	m.Paused = true
	if !strings.Contains(m.View(), "No moves left.") {
		t.Fatal("reopening a lost game should still show the restart prompt")
	}
}

func TestWinningLastMoveStillCelebrates2048(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	m.Board = Board{{1024, 1024, 8, 16}, {8, 16, 8, 16}, {16, 8, 16, 8}, {8, 16, 8, 16}}
	m = press(m, "a")
	if !m.Won || !m.GameOver || !strings.Contains(m.View(), "2048 REACHED!") {
		t.Fatal("reaching 2048 on the last legal move should still show victory")
	}
	m.Paused = true // Returning from the launcher must still allow resume.
	m = press(m, "p")
	if m.Paused {
		t.Fatal("could not resume the winning board")
	}
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	m = updated.(Model)
	if !m.KeepPlaying || !strings.Contains(m.View(), "No moves left.") {
		t.Fatal("continuing a full board should explain there are no moves left")
	}
}

func TestPauseAndAnimationLifecycle(t *testing.T) {
	m := newModel(rand.New(rand.NewSource(1)))
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeySpace})
	m = updated.(Model)
	before := m.Board
	m = press(m, "d")
	if !m.Paused || m.Board != before {
		t.Fatal("Space must pause and block moves")
	}
	m = press(m, "p")
	m.Board = Board{{2, 2}}
	m = press(m, "a")
	oldAnimation := m.animation
	m = press(m, "s")
	updated, _ = m.Update(settleMsg(oldAnimation))
	m = updated.(Model)
	if !m.flashing {
		t.Fatal("stale animation message cleared the newer highlight")
	}
	updated, _ = m.Update(settleMsg(m.animation))
	if updated.(Model).flashing {
		t.Fatal("current animation did not settle")
	}
	m = press(m, "r")
	updated, _ = m.Update(settleMsg(oldAnimation))
	if updated.(Model).Board != m.Board {
		t.Fatal("stale animation changed a restarted game")
	}
}

func TestViewFitsStandardTerminalWithColors(t *testing.T) {
	previous := lipgloss.ColorProfile()
	t.Cleanup(func() { lipgloss.SetColorProfile(previous) })
	for _, profile := range []termenv.Profile{termenv.TrueColor, termenv.ANSI256, termenv.Ascii} {
		lipgloss.SetColorProfile(profile)
		for _, state := range []string{"playing", "paused", "won", "over"} {
			m := newModel(rand.New(rand.NewSource(1)))
			m.Paused, m.Won, m.GameOver = state == "paused", state == "won", state == "over"
			m.Board = Board{{2, 4, 8, 16}, {32, 64, 128, 256}, {512, 1024, 2048, 4096}}
			if !ui.Fits(80, 24, m.View()) {
				t.Fatalf("%s view in %s exceeds 80x24", state, profile.Name())
			}
		}
	}
}
