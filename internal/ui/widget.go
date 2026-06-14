package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
)



func Line(rowData [9]int, cursor [2]int, line int, cursorValue int) string {
	var cells []string
	var cursorSelected, editable, sameColLine, sameValue, mistake, isTheMistakeCell bool
	var mistakeIndex int
	for i := range len(rowData) {
		if rowData[i] > 9 {
			mistake = true
			mistakeIndex = i
			break
		}
		mistake = false
	}
	for i, v := range rowData {
		var value string
		cellsPostion := [2]int{i, line}
		cursorSelected = cursor[0] == i && cursor[1] == line
		editable = logic.IsEditable(cellsPostion)
		sameColLine = cursor[0] == i || cursor[1] == line
		sameValue = v == cursorValue && v != 0 



		// handling if cell is empty or not
		if v == 0 {
			value = " "
		}else if logic.IsNotValid(line, i, v) {
			value = fmt.Sprint(v - 207)
		}else{
			value = fmt.Sprint(v)
		}

		

		// cursor and editable styling handling
		if i == mistakeIndex {
			cells = append(cells, CellStyle(cursorSelected, editable, sameColLine, sameValue, mistake, isTheMistakeCell).Render(value))
		}else{
			cells = append(cells, CellStyle(cursorSelected, editable, sameColLine, sameValue, mistake, isTheMistakeCell).Render(value))
		}
		
		// Add separator after every 3 cells except the last block
		if (i+1)%3 == 0 && i != 8 {
			cells = append(cells, "  ")
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}