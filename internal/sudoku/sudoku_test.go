package sudoku

import (
	tea "github.com/charmbracelet/bubbletea"
	"testing"
	"time"
)

func TestCluesAndPausedCellsAreProtected(t *testing.T) {
	m := NewModel()
	m.SelectingLevel = false
	m.Puzzle[0][0], m.Cells[0][0] = 5, 5
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyBackspace})
	m = updated.(Model)
	if m.Cells[0][0] != 5 {
		t.Fatal("erased a given clue")
	}
	m.Puzzle[0][0], m.Cells[0][0], m.Solution[0][0] = 0, 0, 5
	m.Paused = true
	updated, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("5")})
	if updated.(Model).Cells[0][0] != 0 {
		t.Fatal("edited while paused")
	}
}

func TestClockOnlyRunsDuringPlay(t *testing.T) {
	for _, m := range []Model{{SelectingLevel: true}, {Paused: true}, {Restarting: true}, {Won: true}, {GameOver: true}} {
		updated, _ := m.Update(tickMsg(time.Now()))
		if updated.(Model).Elapsed != 0 {
			t.Fatal("inactive clock advanced")
		}
	}
	updated, _ := (Model{}).Update(tickMsg(time.Now()))
	if updated.(Model).Elapsed != time.Second {
		t.Fatal("active clock did not advance")
	}
}
