package game2048

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
)

func (m Model) View() string {
	var rows []string
	for row := range m.Board {
		var tiles []string
		for column, value := range m.Board[row] {
			label := "·"
			if value != 0 {
				label = fmt.Sprint(value)
			}
			style := tileStyle(value)
			if m.flashing && m.merged[row][column] {
				style = style.Background(ui.Fixed).Foreground(ui.Surface)
			} else if m.flashing && m.spawned == [2]int{row, column} {
				label = "·" + label + "·"
			}
			tiles = append(tiles, style.Render(label))
			if column < size-1 {
				tiles = append(tiles, " ")
			}
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, tiles...))
		if row < size-1 {
			rows = append(rows, "")
		}
	}
	board := lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(ui.Primary).
		Padding(1).Render(strings.Join(rows, "\n"))

	highest := m.Board.highest()
	target := 2048
	if highest >= target {
		target = highest * 2
	}
	status := ui.State(m.Paused, m.GameOver, false)
	tip := "Merge tiles to reach 2048."
	if m.GameOver {
		tip = "No moves left.\nR to try again."
	} else if m.Paused {
		tip = "P or Space to resume."
	} else if m.Won && !m.KeepPlaying {
		status = lipgloss.NewStyle().Foreground(ui.Success).Bold(true).Render("2048 REACHED!")
		tip = "Enter Keep playing\nR Start a new board"
	} else if m.lastGain > 0 {
		tip = fmt.Sprintf("+%d from your last move", m.lastGain)
	}
	stats := status + "\n\n" + ui.Stat("SCORE", m.Score) + "\n\n" +
		ui.Stat("SESSION BEST", m.BestScore) + "\n\n" +
		ui.Stat("MOVES / HIGHEST", fmt.Sprintf("%d / %d", m.Moves, highest)) + "\n\n" +
		ui.Stat("NEXT TARGET", target) + "\n\n" +
		ui.CmdStyle.Render("↑↓←→ / WASD Slide\nSpace Pause") + "\n\n" + tip
	return ui.Playfield(m.Width, board, ui.Panel("2048 · MERGE & GROW", stats, 26))
}

func tileStyle(value int) lipgloss.Style {
	background, foreground := ui.PanelColor, ui.Muted
	if value != 0 {
		palette := []lipgloss.Color{
			"#34456B", "#365C91", "#23736E", "#337B4C", "#C8A753",
			"#D98462", "#D5773F", "#8A5FC1", "#AC83D4", "#87BFE8", "#F2C94C",
		}
		index := 0
		for tile := value; tile > 2; tile /= 2 {
			index++
		}
		background = palette[min(index, len(palette)-1)]
		foreground = ui.Fixed
		if value >= 32 && value <= 128 || value >= 512 {
			foreground = ui.Surface
		}
	}
	return lipgloss.NewStyle().Width(8).Align(lipgloss.Center).Padding(1, 0).
		Background(background).Foreground(foreground).Bold(value != 0)
}
