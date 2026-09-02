package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
	"github.com/misafidiniaina/sudoku/internal/logic/gen"
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

func GameBoard(Data [9][9]int, puzzle [9][9]int, cursor [2]int) string {
	var result string
	cursorRaw := Data[cursor[1]][cursor[0]]
	cursorIsMistake := cursorRaw < 0

	var cursorValue int
	if !cursorIsMistake {
		cursorValue = logic.GetValueInCursor(Data, cursor)
	} else {
		cursorValue = 0 // disables sameValue highlighting for all cells
	}

	for i := range 9 {
		if i == 2 || i == 5 {
			result = result + Line(Data[i], puzzle, cursor, i, cursorValue) + "\n\n"
		} else {
			result = result + Line(Data[i], puzzle, cursor, i, cursorValue) + "\n"
		}
	}
	return result
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
	var result string

	result = "Move:" + CmdStyle.Render(" ↑ ← ↓ → ") + "       Enter number: " + CmdStyle.Render("1-9 ") + "       Clear cell: " + CmdStyle.Render("Backspace/Delete") + "\nPause/Resume: " + CmdStyle.Render("p") + "      Quit: " + CmdStyle.Render("q")
	return result

}
