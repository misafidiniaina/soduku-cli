package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
)



func Line(rowData [9]int, cursor [2]int, line int) string {
	var cells []string

	for i, v := range rowData {
		var value string

		// handling if cell is empty or not
		if v == 0 {
			value = " "
		} else {
			value = fmt.Sprint(v)		
		}

		// cursor and editable styling handling
		
		if cursor[0] == i && cursor[1] == line {
			if logic.IsEditable(cursor) {
				cells = append(cells, EditableSelectedCellStyle.Render(value))
			}else{
				cells = append(cells, NoEditableSelectedCellStyle.Render(value))
			}
		}else{
			if logic.IsEditable(cursor) {
				cells = append(cells, EditableCellStyle.Render(value))
			}else{
				cells = append(cells, NoEditableCellsStyle.Render(value))
			}
			
		}

		// Add separator after every 3 cells except the last block
		if (i+1)%3 == 0 && i != 8 {
			cells = append(cells, "  ")
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}