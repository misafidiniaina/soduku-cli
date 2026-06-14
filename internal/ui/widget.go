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

		// string formating for value displaying
        var value string
        if v == 0 {
            value = " "
        } else {
            value = fmt.Sprint(displayValue)
        }

        cellPosition := [2]int{i, line}
		// if the current cell match the cursor postion use cursorSelected boolean variable to apply the corresponding style
        cursorSelected := cursor[0] == i && cursor[1] == line

		// if the current cell is editable apply the corresponding style
        editable := logic.IsEditable(cellPosition)

		// apply corresponding style to the same column and row as the cursor
        sameColLine := cursor[0] == i || cursor[1] == line
        
        // normalize both sides before comparing (absolute value)
        absV := v
        if absV < 0 { absV = -absV }
        absCursor := cursorValue
        if absCursor < 0 { absCursor = -absCursor }
		// compare the absolute value in the rowData[i] (v in the for loop) and the absolute value in the cursor
        sameValue := absV == absCursor && absCursor != 0 // absCursor is the key thing to not applying the samevalue hightlighting because it can be change without altering the m.Cells value in the main/GameTable function

        cells = append(cells, CellStyle(cursorSelected, editable, sameColLine, sameValue, isMistake).Render(value))

        if (i+1)%3 == 0 && i != 8 {
            cells = append(cells, "  ")
        }
    }

    return lipgloss.JoinHorizontal(lipgloss.Center, cells...)
}