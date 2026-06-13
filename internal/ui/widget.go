package ui


import (

	"fmt"
	"github.com/charmbracelet/lipgloss"
)



func Line(rowData [9]int, cursor [2]int, line int) string {
	var cells []string

	for i, v := range rowData {
		var value string
		if v == 0 {
			value = " "
		} else {
			value = fmt.Sprint(v)		}
		if cursor[0] == i && cursor[1] == line {
			cells = append(cells, SelectedCellStyle.Render(value))
		}else{
			cells = append(cells, CellsStyle.Render(value))
		}
		

		// Add separator after every 3 cells except the last block
		if (i+1)%3 == 0 && i != 8 {
			cells = append(cells, "  ")
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}