package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/misafidiniaina/soduku-cli/internal/snake"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku"
	"testing"
)

func TestMenuNavigationAndResize(t *testing.T) {
	m := newAppModel()
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyUp})
	m = next.(appModel)
	if m.selected != 2 {
		t.Fatal("up should wrap backwards")
	}
	next, _ = m.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	if next.(appModel).games[0].(sudoku.Model).Width != 100 {
		t.Fatal("menu swallowed resize")
	}
}

func TestReturningToGamePausesWithoutStartingAnotherTimer(t *testing.T) {
	m := newAppModel()
	m.selected = 1
	next, cmd := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd == nil {
		t.Fatal("first launch needs a timer")
	}
	m = next.(appModel)
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("m")})
	m = next.(appModel)
	if m.active != -1 || !m.games[1].(snake.Model).Paused {
		t.Fatal("menu did not pause snake")
	}
	next, cmd = m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd != nil || next.(appModel).active != 1 {
		t.Fatal("resume must reuse the existing timer")
	}
}
