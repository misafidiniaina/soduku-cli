package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
)

func GameBoard(Data [9][9]int, cursor [2]int) string{
	var result string
	cursorValue := logic.GetValueInCursor(Data, cursor)
	for i := range 9 {
		if i == 2 || i == 5 {
			result = result + Line(Data[i], cursor, i, cursorValue) +"\n\n"
		}else{
			result = result + Line(Data[i], cursor, i, cursorValue) +"\n"
		}
		
	}
	return result
}

func GameHeader(score int,error int, time string) string{
	var result string

	scoreItem := HeadTextStyle.Render("Score: ") + fmt.Sprint(score)
	errorItem := HeadTextStyle.Render("   Mistakes: ") + fmt.Sprint(error)
	timeItem := HeadTextStyle.Render("  Time: ") + time


	result = lipgloss.JoinHorizontal(
		lipgloss.Center,
		HeadItemStyle.Render(scoreItem),
		HeadItemStyle.Render(errorItem),
		HeadItemStyle.Render(timeItem),
	)

	return result
}

func CommandHelper() string {
	var result string

	result = "Move:" + CmdStyle.Render(" ↑ ← ↓ → ") +"       Enter number: "+CmdStyle.Render("1-9 ")+"       Clear cell: "+CmdStyle.Render("Backspace/Delete")+"       Quit: "+CmdStyle.Render("q")
	return  result

}