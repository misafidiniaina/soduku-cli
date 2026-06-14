package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
)



func Line(rowData [9]int, cursor [2]int, line int, cursorValue int) string {
    var cells []string

    for i, v := range rowData {
        isMistake := v < 0
        displayValue := v
        if isMistake {
            displayValue = -v
        }

        var value string
        if v == 0 {
            value = " "
        } else {
            value = fmt.Sprint(displayValue)
        }

        cellPosition := [2]int{i, line}
        cursorSelected := cursor[0] == i && cursor[1] == line
        editable := logic.IsEditable(cellPosition)
        sameColLine := cursor[0] == i || cursor[1] == line
        
        // normalize both sides before comparing
        absV := v
        if absV < 0 { absV = -absV }
        absCursor := cursorValue
        if absCursor < 0 { absCursor = -absCursor }
        sameValue := absV == absCursor && absCursor != 0

        cells = append(cells, CellStyle(cursorSelected, editable, sameColLine, sameValue, isMistake).Render(value))

        if (i+1)%3 == 0 && i != 8 {
            cells = append(cells, "  ")
        }
    }

    return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}