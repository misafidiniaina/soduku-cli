package minesweeper

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
)

func (m Model) View() string {
	if m.SelectingDifficulty {
		return m.difficultyView()
	}
	board := m.boardView()
	remaining := m.Board.Mines - m.Board.Flags
	status := ui.State(m.Paused, m.GameOver, m.Won)
	tip := "Reveal every safe cell."
	if m.Paused {
		tip = "P to resume."
	} else if m.GameOver {
		tip = "Mine hit. R to try again."
	} else if m.Won {
		tip = "Board cleared! R for a new board."
	}
	stats := status + "\n\n" +
		ui.Stat("MINES LEFT", remaining) + "\n\n" +
		ui.Stat("TIME", formatTime(m.Elapsed)) + "\n\n" +
		ui.Stat("SAFE CELLS", fmt.Sprintf("%d / %d", m.Board.Revealed, m.Board.Width*m.Board.Height-m.Board.Mines)) + "\n\n" +
		ui.CmdStyle.Render("↑↓←→ / WASD Move\nEnter / Space Reveal\nF Flag · P Pause") + "\n\n" + tip
	return ui.Playfield(m.Width, board, ui.Panel("MINESWEEPER", stats, 26))
}

func (m Model) difficultyView() string {
	var options []string
	for index, difficulty := range difficulties {
		label := fmt.Sprintf("%d. %-12s %2d×%-2d · %d mines", index+1, difficulty.Name, difficulty.Width, difficulty.Height, difficulty.Mines)
		style := ui.LevelOptionStyle
		if index == m.DifficultyIndex {
			label = "› " + label
			style = ui.LevelSelectedStyle
		} else {
			label = "  " + label
		}
		options = append(options, style.Render(label))
	}
	content := ui.LevelTitleStyle.Render("MINESWEEPER · SELECT DIFFICULTY") + "\n\n" +
		strings.Join(options, "\n") + "\n\n" +
		ui.CmdStyle.Render("Use arrows or 1–3, then press Enter")
	return ui.WrapperStyle.MarginTop(0).Render(content)
}

func (m Model) boardView() string {
	rows := make([]string, 0, m.Board.Height)
	cellWidth := 2
	if m.Board.Width > 20 {
		cellWidth = 1
	}
	for row := 0; row < m.Board.Height; row++ {
		cells := make([]string, 0, m.Board.Width)
		for col := 0; col < m.Board.Width; col++ {
			point := Point{Row: row, Col: col}
			cell := m.Board.Cells[row][col]
			label := "■"
			style := lipgloss.NewStyle().Foreground(ui.Muted).Background(ui.PanelColor)
			switch {
			case cell.Revealed && cell.Mine:
				label = "*"
				style = lipgloss.NewStyle().Foreground(ui.Fixed).Background(ui.Secondary).Bold(true)
			case cell.Revealed:
				label = " "
				style = lipgloss.NewStyle().Foreground(ui.Muted).Background(ui.Surface)
				if cell.Adjacent > 0 {
					label = fmt.Sprint(cell.Adjacent)
					style = numberStyle(cell.Adjacent)
				}
			case cell.Flagged:
				label = "F"
				style = lipgloss.NewStyle().Foreground(ui.Warning).Background(ui.PanelColor).Bold(true)
			}
			if point == m.Cursor {
				style = style.Background(ui.Warning).Foreground(ui.Surface).Bold(true)
			}
			cells = append(cells, style.Width(cellWidth).Align(lipgloss.Center).Render(label))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cells...))
	}
	return lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(ui.Primary).
		Padding(0, 1).Render(strings.Join(rows, "\n"))
}

func numberStyle(number int) lipgloss.Style {
	colors := []lipgloss.Color{
		ui.Muted,
		"#6AA9FF",
		"#59C98C",
		"#FF6B6B",
		"#B69CFF",
		"#F2A900",
		"#45D5CF",
		"#D98462",
		ui.Fixed,
	}
	return lipgloss.NewStyle().Foreground(colors[number]).Background(ui.Surface).Bold(true)
}

func formatTime(elapsed time.Duration) string {
	seconds := int(elapsed.Seconds())
	return fmt.Sprintf("%02d:%02d", seconds/60, seconds%60)
}
