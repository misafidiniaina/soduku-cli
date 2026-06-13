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
	Initial	  = lipgloss.Color("#4b69ff")
)


var(
	NoEditableCellsStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(Muted).
		Align(lipgloss.Center).
		PaddingLeft(1).
		PaddingRight(1).
		MarginRight(1).
		Bold(true)

	NoEditableSelectedCellStyle = NoEditableCellsStyle.
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Secondary)
	
	EditableCellStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(Muted).
		Align(lipgloss.Center).
		PaddingLeft(1).
		PaddingRight(1).
		MarginRight(1)
	
	EditableSelectedCellStyle = EditableCellStyle.
		Border(lipgloss.DoubleBorder()).
		BorderForeground(Secondary)
		

	WelcomeStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff0000")).
		Border(lipgloss.NormalBorder())

	FooterStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00ff00")).
		Border(lipgloss.NormalBorder())


	BigblockStyle = lipgloss.NewStyle().
		Width(60).
		Border(lipgloss.NormalBorder()).
		BorderBackground(lipgloss.Color("#0dfsdf")).
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("#000")).
		Bold(true)

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