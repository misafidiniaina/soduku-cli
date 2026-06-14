package ui

import (
	"github.com/charmbracelet/lipgloss"
)


var(
	Primary   = lipgloss.Color("#7D56F4")
	Secondary = lipgloss.Color("#ff005d")
	Success   = lipgloss.Color("#04B575")
	Warning   = lipgloss.Color("#F2A900")
	Muted     = lipgloss.Color("#a1a1a1")
	Command   = lipgloss.Color("#87CEEB")
	Editable  = lipgloss.Color("#5f7aff")
	Fixed 	  = lipgloss.Color("#ffffff")
)


var(

	// cell base style
	BaseCell = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Align(lipgloss.Center).
		PaddingLeft(1).
		PaddingRight(1).
		MarginRight(1)

	// same value as cursor selected
	WrapperStyle = lipgloss.NewStyle().
		MarginLeft(5).
		MarginTop(1)

	HeadItemStyle = lipgloss.NewStyle().
		Width(18)

	HeadTextStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(Success))


	CmdStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Command))
)


func CellStyle(cursorSelected, editable, inSameCol_L, sameValue, isMistake bool) lipgloss.Style {
    style := BaseCell

	// The key thing here is to make a condition about same styling purpose (ex: condition for border color)
	// if the same topic is changed in different the last one will be applied (ex here: the borderforeground is applied if the condition true)

    if inSameCol_L {
        style = style.BorderForeground(Fixed)
    } else {
        style = style.BorderForeground(Muted)
    }

    if sameValue && !isMistake {
        style = style.Background(Warning)
    }

    if isMistake {
        style = style.Foreground(Secondary).Bold(true)
    } else if editable {
        style = style.Foreground(Editable).Bold(false)
    } else {
        style = style.Foreground(Fixed).Bold(true)
    }

    if cursorSelected {
        style = style.Border(lipgloss.DoubleBorder()).BorderForeground(Secondary)
    }

    return style
}