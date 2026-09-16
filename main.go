package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/snake"
	"github.com/misafidiniaina/sudoku/internal/sudoku"
	"github.com/misafidiniaina/sudoku/internal/ui"
)

type appModel struct {
	games    []tea.Model
	names    []string
	selected int
	active   int
}

func newAppModel() appModel {
	return appModel{
		games:  []tea.Model{sudoku.NewModel(), snake.NewModel()},
		names:  []string{"Sudoku", "Snake"},
		active: -1,
	}
}

func (m appModel) Init() tea.Cmd {
	return nil
}

func (m appModel) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if m.active < 0 {
		if keyMessage, ok := message.(tea.KeyMsg); ok {
			switch keyMessage.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "up", "left", "down", "right":
				m.selected = (m.selected + 1) % len(m.games)
			case "enter":
				m.active = m.selected
				return m, m.games[m.active].Init()
			}
		}
		return m, nil
	}

	if keyMessage, ok := message.(tea.KeyMsg); ok && keyMessage.String() == "m" {
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

	options := make([]string, len(m.names))
	for index, name := range m.names {
		style := lipgloss.NewStyle().Padding(0, 2)
		if index == m.selected {
			style = style.Bold(true).Foreground(ui.Success).Background(ui.Primary)
		}
		options[index] = style.Render(name)
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

func main() {
	p := tea.NewProgram(newAppModel())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
