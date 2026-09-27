package ui

import (
	"github.com/charmbracelet/lipgloss"
	"strings"
	"testing"
)

func TestScreenDimensions(t *testing.T) {
	for _, size := range [][2]int{{80, 24}, {120, 40}, {40, 12}, {10, 3}, {1, 1}, {1, 2}} {
		for _, content := range []string{"board", strings.Repeat("X", 100) + "\n" + strings.Repeat("row\n", 40)} {
			view := Screen(size[0], size[1], "SUDOKU", content, "P Pause · M Menu · Q Quit")
			if lipgloss.Width(view) != size[0] || lipgloss.Height(view) != size[1] {
				t.Fatalf("%v: got %dx%d", size, lipgloss.Width(view), lipgloss.Height(view))
			}
		}
	}
}
