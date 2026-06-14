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


func CellStyle(cursorSelected, editable, inSameCol_L, sameValue bool) lipgloss.Style {
	style := BaseCell

	if inSameCol_L {
		style = style.BorderForeground(Fixed)
	}else{
		style = style.BorderForeground(Muted)
	}

	if sameValue {
		style = style.Background(nil)
	}else{
		style = style.Background(nil)
	}


	if editable {
		style = style.Foreground(Editable).Bold(false)
	}else{
		style = style.Foreground(Fixed).Bold(true)
	}

	
	
	if cursorSelected {
		style = style.Border(lipgloss.DoubleBorder()).BorderForeground(Secondary)
	}

	return style
}