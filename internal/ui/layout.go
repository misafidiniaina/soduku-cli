package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	Surface    = lipgloss.Color("#111522")
	PanelColor = lipgloss.Color("#1C2233")
)

// Screen keeps the navigation at the edges and centers the play area in the
// remaining space. The returned view always occupies exactly the terminal.
func Screen(width, height int, title, content, controls string) string {
	width, height = max(1, width), max(1, height)
	bar := lipgloss.NewStyle().Background(PanelColor).Foreground(Fixed).Width(width).MaxWidth(width).MaxHeight(1)
	header := bar.Bold(true).Render("  CLI GAMES  /  " + title)
	footer := bar.Foreground(Command).Render("  " + controls)
	if height == 1 {
		return header
	}
	bodyHeight := height - 2
	if !Fits(width, height, content) {
		content = LevelTitleStyle.Render("MORE ROOM TO PLAY") + "\n\n" +
			fmt.Sprintf("Resize to at least %d × %d", lipgloss.Width(content), lipgloss.Height(content)+2) +
			"\nYour game is paused. Press P after resizing.\nM Menu · Q Quit"
	}
	content = lipgloss.NewStyle().Background(Surface).Foreground(Fixed).MaxWidth(width).MaxHeight(max(0, bodyHeight)).Render(content)
	body := lipgloss.Place(width, bodyHeight, lipgloss.Center, lipgloss.Center, content,
		lipgloss.WithWhitespaceBackground(Surface))
	if bodyHeight == 0 {
		return header + "\n" + footer
	}
	return header + "\n" + paintSurface(body) + "\n" + footer
}

func Fits(width, height int, content string) bool {
	return lipgloss.Width(content) <= width && lipgloss.Height(content) <= height-2
}

// Playfield aligns the board and a fixed-width information column. Narrow
// terminals stack the same information below the board.
func Playfield(width int, board, panel string) string {
	board = strings.TrimRight(board, "\n")
	if width <= 0 {
		width = 80
	}
	if lipgloss.Width(board)+2+lipgloss.Width(panel) <= width {
		return lipgloss.JoinHorizontal(lipgloss.Top, board, "  ", panel)
	}
	return lipgloss.JoinVertical(lipgloss.Left, board, "", panel)
}

func Panel(title, content string, width int) string {
	heading := LevelTitleStyle.Render(title)
	return lipgloss.NewStyle().Width(width).Render(heading + "\n" + content)
}

func Stat(label string, value any) string {
	return lipgloss.NewStyle().Foreground(Muted).Render(label) + "\n" +
		HeaderValueStyle.Render(fmt.Sprint(value))
}

func State(paused, over, won bool) string {
	if won {
		return lipgloss.NewStyle().Foreground(Success).Bold(true).Render("BOARD COMPLETE")
	}
	if over {
		return lipgloss.NewStyle().Foreground(Secondary).Bold(true).Render("GAME OVER")
	}
	if paused {
		return lipgloss.NewStyle().Foreground(Warning).Bold(true).Render("PAUSED")
	}
	return lipgloss.NewStyle().Foreground(Success).Render("● PLAYING")
}

// Lip Gloss's nested styles emit full ANSI resets. Restore the canvas colors
// after those resets so gaps, borders, and plain labels never fall back to the
// user's terminal theme. Explicit cell and selection colors still take priority.
func paintSurface(content string) string {
	profile := lipgloss.ColorProfile()
	fg := profile.Color(string(Fixed)).Sequence(false)
	bg := profile.Color(string(Surface)).Sequence(true)
	if fg == "" || bg == "" {
		return content
	}
	base := "\x1b[" + fg + ";" + bg + "m"
	restore := strings.NewReplacer(
		"\x1b[0m", "\x1b[0m"+base,
		"\x1b[m", "\x1b[m"+base,
		"\x1b[39m", "\x1b["+fg+"m",
		"\x1b[49m", "\x1b["+bg+"m",
	)
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		// Rows must be independent for Bubble Tea's partial screen redraws.
		lines[i] = base + restore.Replace(line) + "\x1b[0m"
	}
	return strings.Join(lines, "\n")
}
