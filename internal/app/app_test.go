package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/snake"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku"
	"strings"
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

func TestFullScreenViewsAcrossWindowSizes(t *testing.T) {
	for _, size := range [][2]int{{80, 24}, {80, 31}, {120, 40}, {160, 50}, {35, 10}} {
		for active := -1; active < 3; active++ {
			m := newAppModel()
			next, _ := m.Update(tea.WindowSizeMsg{Width: size[0], Height: size[1]})
			m = next.(appModel)
			m.active = active
			if active == 0 {
				game := m.games[0].(sudoku.Model)
				game.SelectingLevel = false
				m.games[0] = game
			}
			view := m.View()
			if lipgloss.Width(view) != size[0] || lipgloss.Height(view) != size[1] {
				t.Fatalf("game %d at %v: got %dx%d", active, size, lipgloss.Width(view), lipgloss.Height(view))
			}
			if size[0] >= 80 && size[1] >= 31 && strings.Contains(view, "MORE ROOM") {
				t.Fatalf("game %d unnecessarily hidden at %v", active, size)
			}
		}
	}
}

func TestSmallWindowPausesAndBlocksHiddenGameplay(t *testing.T) {
	m := newAppModel()
	m.active = 1
	next, _ := m.Update(tea.WindowSizeMsg{Width: 20, Height: 10})
	m = next.(appModel)
	if !m.games[1].(snake.Model).Paused {
		t.Fatal("small window did not pause")
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	m = next.(appModel)
	if !m.games[1].(snake.Model).Paused {
		t.Fatal("resumed invisible gameplay")
	}
	next, _ = m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	m = next.(appModel)
	if !m.games[1].(snake.Model).Paused {
		t.Fatal("resizing silently resumed gameplay")
	}
	next, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("p")})
	if next.(appModel).games[1].(snake.Model).Paused {
		t.Fatal("could not resume after resizing")
	}
}
