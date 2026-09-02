package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
	"github.com/misafidiniaina/sudoku/internal/logic/gen"
)

func LevelSelector(levels []gen.Difficulty, selected int) string {
	items := "Select a difficulty\n\n"
	for i, level := range levels {
		prefix := "  "
		if i == selected {
			prefix = "> "
		}
		items += fmt.Sprintf("%s%d. %s\n", prefix, i+1, level)
	}
	return HeadItemStyle.Render(items) + "\nUse arrows or 1-6, then press Enter"
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

	scoreItem := HeadTextStyle.Render("Score: ") + fmt.Sprint(score) + "\n"
	errorItem := HeadTextStyle.Render("   Mistakes: ") + fmt.Sprint(error) + "\n"
	timeItem := HeadTextStyle.Render("  Time: ") + time + "\n"
	levelItem := HeadTextStyle.Render("Level: ") + level + "\n"

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
