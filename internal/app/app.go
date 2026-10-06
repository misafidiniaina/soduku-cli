package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/game2048"
	"github.com/misafidiniaina/soduku-cli/internal/minesweeper"
	"github.com/misafidiniaina/soduku-cli/internal/snake"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku"
	"github.com/misafidiniaina/soduku-cli/internal/tetris"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
)

type appModel struct {
	games         []tea.Model
	names         []string
	selected      int
	active        int
	started       []bool
	width, height int
}

func newAppModel() appModel {
	games := []tea.Model{sudoku.NewModel(), snake.NewModel(), tetris.NewModel(), game2048.NewModel(), minesweeper.NewModel()}
	return appModel{
		games:   games,
		names:   []string{"Sudoku", "Snake", "Tetris", "2048", "Minesweeper"},
		started: make([]bool, len(games)),
		active:  -1,
		width:   80, height: 24,
	}
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if size, ok := message.(tea.WindowSizeMsg); ok {
		m.width, m.height = max(1, size.Width), max(1, size.Height)
	}
	if m.active >= 0 && !ui.Fits(m.width, m.height, m.games[m.active].View()) {
		m.pauseActive()
	}
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
		m.pauseActive()
		m.active = -1
		return m, nil
	}

	if key, ok := message.(tea.KeyMsg); ok && key.String() != "q" && key.String() != "ctrl+c" &&
		!ui.Fits(m.width, m.height, m.games[m.active].View()) {
		return m, nil
	}
	updated, command := m.games[m.active].Update(message)
	m.games[m.active] = updated
	return m, command
}

func (m appModel) View() string {
	if m.active >= 0 {
		controls := "P Pause / Resume   R Restart   M Menu   Q Quit"
		if game, ok := m.games[m.active].(sudoku.Model); ok {
			if game.SelectingLevel {
				controls = "↑↓ / 1–6 Choose   Enter Confirm   M Menu   Q Quit"
			}
			if game.Restarting {
				controls = "↑↓ Choose   Enter Confirm   Esc Cancel   M Menu   Q Quit"
			}
		} else if game, ok := m.games[m.active].(minesweeper.Model); ok {
			if game.SelectingDifficulty {
				controls = "↑↓ Choose   1–3 Select   Enter Start   M Menu   Q Quit"
			} else {
				controls = "↑↓←→ / WASD Move   Enter Reveal   F Flag   P Pause   R Restart   M Menu   Q Quit"
			}
		}
		return ui.Screen(m.width, m.height, m.names[m.active], m.games[m.active].View(), controls)
	}

	descriptions := []string{"Find the pattern · Six levels", "Chase the food · Beat your best", "Stack and clear · Plan your next drop", "Slide and merge · Reach 2048", "Reveal safe cells · Flag mines"}
	options := make([]string, len(m.names))
	for index, name := range m.names {
		style := lipgloss.NewStyle().Padding(1, 1).Width(36)
		prefix := "  "
		if index == m.selected {
			prefix = "› "
			style = style.Bold(true).Foreground(ui.Fixed).Background(ui.Primary)
		}
		options[index] = style.Render(prefix + name + "\n  " + descriptions[index])
	}
	var rows []string
	for index := 0; index < len(options); index += 2 {
		if index+1 < len(options) {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, options[index], "  ", options[index+1]))
		} else {
			rows = append(rows, options[index])
		}
	}

	menu := lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.NewStyle().Bold(true).Foreground(ui.Primary).Render("CLI GAMES"),
		"",
		lipgloss.JoinVertical(lipgloss.Left, rows...),
		"",
		ui.CmdStyle.Render("[↑/↓] Choose   [Enter] Start   [Q] Quit"),
	)
	return ui.Screen(m.width, m.height, "ARCADE", menu, "↑↓ Choose   Enter Play   Q Quit")
}

func (m *appModel) pauseActive() {
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
	case game2048.Model:
		game.Paused = true
		m.games[m.active] = game
	case minesweeper.Model:
		game.Paused = true
		m.games[m.active] = game
	}
}

func Run() error {
	p := tea.NewProgram(newAppModel(), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
