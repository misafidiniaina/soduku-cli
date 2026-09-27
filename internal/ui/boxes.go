package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku/logic/gen"
)

func LevelSelector(levels []gen.Difficulty, selected int) string {
	items := LevelTitleStyle.Render("SUDOKU • SELECT DIFFICULTY") + "\n\n"
	for i, level := range levels {
		prefix := "  "
		style := LevelOptionStyle
		if i == selected {
			prefix = "> "
			style = LevelSelectedStyle
		}
		items += style.Render(fmt.Sprintf("%s%d. %s", prefix, i+1, level)) + "\n"
	}
	return items + CmdStyle.Render("Use arrows or 1–6, then press Enter")
}

func RestartSelector(selected int) string {
	options := []string{"Replay existing board", "Generate a new board", "Choose another level"}
	result := LevelTitleStyle.Render("RESTART GAME") + "\n\n"
	for i, option := range options {
		prefix, style := "  ", LevelOptionStyle
		if i == selected {
			prefix, style = "> ", LevelSelectedStyle
		}
		result += style.Render(prefix+option) + "\n"
	}
	return result + "\n" + CmdStyle.Render("Use arrows, then press Enter")
}

func GameBoard(Data [9][9]int, puzzle [9][9]int, cursor [2]int, width int) string {
	var rows []string
	divider := "├─────────┼─────────┼─────────┤"
	rows = append(rows, "┌─────────┬─────────┬─────────┐")
	for y := 0; y < 9; y++ {
		line := "│"
		for x := 0; x < 9; x++ {
			v := Data[y][x]
			label := " · "
			if v != 0 {
				label = fmt.Sprintf(" %d ", abs(v))
			}
			style := lipgloss.NewStyle().Foreground(Muted)
			if puzzle[y][x] != 0 {
				style = style.Foreground(Fixed).Bold(true)
			} else if v > 0 {
				style = style.Foreground(Editable)
			}
			if x == cursor[0] || y == cursor[1] || (x/3 == cursor[0]/3 && y/3 == cursor[1]/3) {
				style = style.Background(lipgloss.Color("#202238"))
			}
			if v > 0 && v == Data[cursor[1]][cursor[0]] {
				style = style.Foreground(Warning).Bold(true)
			}
			if v < 0 {
				style = style.Foreground(Secondary)
			}
			if x == cursor[0] && y == cursor[1] {
				style = style.Background(Primary).Foreground(Fixed).Bold(true)
			}
			line += style.Render(label)
			if x%3 == 2 {
				line += "│"
			}
		}
		rows = append(rows, line)
		if y == 2 || y == 5 {
			rows = append(rows, divider)
		}
	}
	rows = append(rows, "└─────────┴─────────┴─────────┘")
	return strings.Join(rows, "\n")
}

func GameHeader(score int, level string, error int, time string) string {
	var result string

	scoreItem := HeadTextStyle.Render("Score: ") + HeaderValueStyle.Render(fmt.Sprint(score)) + "\n"
	errorItem := HeadTextStyle.Render("Mistakes: ") + HeaderValueStyle.Render(fmt.Sprintf("%d/3", error)) + "\n"
	timeItem := HeadTextStyle.Render("Time: ") + HeaderValueStyle.Render(time) + "\n"
	levelItem := HeadTextStyle.Render("Level: ") + HeaderValueStyle.Render(level) + "\n"

	result = lipgloss.JoinHorizontal(
		lipgloss.Center,
		HeadItemStyle.Render(scoreItem),
		HeadItemStyle.Render(errorItem),
		HeadItemStyle.Render(timeItem),
		HeadItemStyle.Render(levelItem),
	)

	return result
}

func CommandHelper() string {
	return CmdStyle.Render("↑↓←→ Move · 1–9 Enter · Backspace Clear\nP Pause · R Restart · M Menu · Q Quit")
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
