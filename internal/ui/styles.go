package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Primary   = lipgloss.Color("#7D56F4")
	Secondary = lipgloss.Color("#ff005d")
	Success   = lipgloss.Color("#04B575")
	Warning   = lipgloss.Color("#F2A900")
	Muted     = lipgloss.Color("#a1a1a1")
	Command   = lipgloss.Color("#87CEEB")
	Editable  = lipgloss.Color("#5f7aff")
	Fixed     = lipgloss.Color("#ffffff")
	Message   = lipgloss.Color("62")
)

var (

	// cell base style
	BaseCell = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			Align(lipgloss.Center).
			PaddingLeft(1).
			PaddingRight(1).
			MarginRight(1)

	// same value as cursor selected
	WrapperStyle = lipgloss.NewStyle().
			MarginTop(1).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(Primary)

	HeadItemStyle = lipgloss.NewStyle().
			Width(18)

	HeadTextStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(Success))
	HeaderValueStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#ffffff")).
				Bold(true)

	CmdStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Command))
	CommandItemStyle = lipgloss.NewStyle().Width(25)

	PausedGameSyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Message).
			Padding(14)
	GameOverStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(Secondary).
			Align(lipgloss.Center).
			Padding(8)
	WinStyle           = GameOverStyle.Copy().Foreground(Success)
	LevelTitleStyle    = lipgloss.NewStyle().Bold(true).Foreground(Primary)
	LevelSelectedStyle = lipgloss.NewStyle().Bold(true).Foreground(Warning)
	LevelOptionStyle   = lipgloss.NewStyle().Foreground(Fixed)
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
