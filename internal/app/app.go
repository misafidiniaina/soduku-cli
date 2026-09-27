package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/snake"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku"
	"github.com/misafidiniaina/soduku-cli/internal/tetris"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
)

type appModel struct {
	games    []tea.Model
	names    []string
	selected int
	active   int
	started  [3]bool
}

func newAppModel() appModel {
	return appModel{
		games:  []tea.Model{sudoku.NewModel(), snake.NewModel(), tetris.NewModel()},
		names:  []string{"Sudoku", "Snake", "Tetris"},
		active: -1,
	}
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if _, isKey := message.(tea.KeyMsg); !isKey {
		var commands []tea.Cmd
		for i, game := range m.games {
			updated, cmd := game.Update(message)
			m.games[i] = updated
			if cmd != nil {
				commands = append(commands, cmd)
			}
		}
		return m, tea.Batch(commands...)
	}
	if m.active < 0 {
		if keyMessage, ok := message.(tea.KeyMsg); ok {
			switch keyMessage.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "up", "left":
				m.selected = (m.selected + len(m.games) - 1) % len(m.games)
			case "down", "right":
				m.selected = (m.selected + 1) % len(m.games)
			case "enter":
				m.active = m.selected
				if !m.started[m.active] {
					m.started[m.active] = true
					return m, m.games[m.active].Init()
				}
				return m, nil
			}
		}
		return m, nil
	}

	if keyMessage, ok := message.(tea.KeyMsg); ok && keyMessage.String() == "m" {
		switch game := m.games[m.active].(type) {
		case sudoku.Model:
			if !game.SelectingLevel && !game.GameOver && !game.Won {
				game.Paused = true
			}
			m.games[m.active] = game
		case snake.Model:
			game.Paused = true
			m.games[m.active] = game
		case tetris.Model:
			game.Paused = true
			m.games[m.active] = game
		}
		m.active = -1
		return m, nil
	}

	updated, command := m.games[m.active].Update(message)
	m.games[m.active] = updated
	return m, command
}

func (m appModel) View() string {
	if m.active >= 0 {
		return m.games[m.active].View()
	}

	descriptions := []string{"Find the pattern · Six difficulty levels", "Chase the food · Beat your best score", "Stack and clear · Plan your next drop"}
	options := make([]string, len(m.names))
	for index, name := range m.names {
		style := lipgloss.NewStyle().Padding(1, 2)
		prefix := "  "
		if index == m.selected {
			prefix = "› "
			style = style.Bold(true).Foreground(ui.Fixed).Background(ui.Primary)
		}
		options[index] = style.Render(prefix + name + "\n  " + descriptions[index])
	}

	menu := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Bold(true).Foreground(ui.Primary).Render("CLI GAMES"),
		"",
		lipgloss.JoinVertical(lipgloss.Left, options...),
		"",
		ui.CmdStyle.Render("[↑/↓] Choose   [Enter] Start   [Q] Quit"),
	)
	return lipgloss.NewStyle().Padding(2, 4).Border(lipgloss.RoundedBorder()).BorderForeground(ui.Primary).Render(menu)
}

func Run() error {
	p := tea.NewProgram(newAppModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
