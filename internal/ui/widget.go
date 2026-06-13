package ui


import (

	"fmt"
	"github.com/charmbracelet/lipgloss"
)



func Line(RowData [9]int) string  {

	cellsStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Align(lipgloss.Center).
		PaddingLeft(1).
		PaddingRight(1).
		MarginRight(1)

	line := lipgloss.JoinHorizontal(
		lipgloss.Center,
		cellsStyle.Render(fmt.Sprint(RowData[0])),
		cellsStyle.Render(fmt.Sprint(RowData[1])),
		cellsStyle.Render(fmt.Sprint(RowData[2])),
		"  ",
		cellsStyle.Render(fmt.Sprint(RowData[3])),
		cellsStyle.Render(fmt.Sprint(RowData[4])),	
		cellsStyle.Render(fmt.Sprint(RowData[5])),
		"  ",
		cellsStyle.Render(fmt.Sprint(RowData[6])),
		cellsStyle.Render(fmt.Sprint(RowData[7])),
		cellsStyle.Render(fmt.Sprint(RowData[8])),
	)
	return line
}