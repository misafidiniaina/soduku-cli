package ui


import (

	"fmt"
	"github.com/charmbracelet/lipgloss"
)



func Line(RowData [9]int) string  {

	line := lipgloss.JoinHorizontal(
		lipgloss.Center,
		CellsStyle.Render(fmt.Sprint(RowData[0])),
		CellsStyle.Render(fmt.Sprint(RowData[1])),
		CellsStyle.Render(fmt.Sprint(RowData[2])),
		"  ",
		CellsStyle.Render(fmt.Sprint(RowData[3])),
		CellsStyle.Render(fmt.Sprint(RowData[4])),	
		CellsStyle.Render(fmt.Sprint(RowData[5])),
		"  ",
		CellsStyle.Render(fmt.Sprint(RowData[6])),
		CellsStyle.Render(fmt.Sprint(RowData[7])),
		CellsStyle.Render(fmt.Sprint(RowData[8])),
	)
	return line
}