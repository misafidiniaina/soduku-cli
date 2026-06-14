package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
)



func Line(rowData [9]int, cursor [2]int, line int) string {
	var cells []string
	var cursorSelected, editable bool
	

	for i, v := range rowData {
		var value string
		cellsPostion := [2]int{i, line}
		cursorSelected = cursor[0] == i && cursor[1] == line
		editable = logic.IsEditable(cellsPostion)

		// handling if cell is empty or not
		if v == 0 {
			value = " "
		} else {
			value = fmt.Sprint(v)		
		}

		// cursor and editable styling handling
		cells = append(cells, CellStyle(cursorSelected, editable).Render(value))

		// Add separator after every 3 cells except the last block
		if (i+1)%3 == 0 && i != 8 {
			cells = append(cells, "  ")
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}