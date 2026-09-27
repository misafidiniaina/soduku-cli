package ui

import (
	"strconv"
	"strings"
	"testing"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/muesli/termenv"
)

// Read the effective color of every printed rune, including spaces. Checking
// ANSI tokens alone misses resets that expose the user's terminal background.
func readColors(text string, visit func(rune, string, string)) {
	fg, bg := "", ""
	parser := ansi.NewParser()
	parser.SetHandler(ansi.Handler{
		Print: func(r rune) { visit(r, fg, bg) },
		HandleCsi: func(cmd ansi.Cmd, params ansi.Params) {
			if cmd.Final() != 'm' {
				return
			}
			if len(params) == 0 {
				fg, bg = "", ""
				return
			}
			for i := 0; i < len(params); i++ {
				code := params[i].Param(0)
				switch {
				case code == 0:
					fg, bg = "", ""
				case code == 39:
					fg = ""
				case code == 49:
					bg = ""
				case code == 38 || code == 48:
					mode, _, _ := params.Param(i+1, 0)
					count := 3
					if mode == 2 {
						count = 5
					}
					var values []string
					for j := 0; j < count && i+j < len(params); j++ {
						values = append(values, strconv.Itoa(params[i+j].Param(0)))
					}
					if code == 38 {
						fg = strings.Join(values, ";")
					} else {
						bg = strings.Join(values, ";")
					}
					i += count - 1
				case code >= 30 && code <= 37 || code >= 90 && code <= 97:
					fg = strconv.Itoa(code)
				case code >= 40 && code <= 47 || code >= 100 && code <= 107:
					bg = strconv.Itoa(code)
				}
			}
		},
	})
	parser.Parse([]byte(text))
}

func TestScreenKeepsColorsAfterNestedResets(t *testing.T) {
	previous := lipgloss.ColorProfile()
	t.Cleanup(func() { lipgloss.SetColorProfile(previous) })
	for _, profile := range []termenv.Profile{termenv.TrueColor, termenv.ANSI256, termenv.ANSI} {
		t.Run(profile.Name(), func(t *testing.T) {
			lipgloss.SetColorProfile(profile)
			colored := lipgloss.NewStyle().Foreground(Secondary).Render("X")
			selected := lipgloss.NewStyle().Background(Primary).Render("Y")
			view := Screen(40, 8, "TEST", "before "+colored+" after "+selected+" end", "controls")
			lines := strings.Split(view, "\n")
			// Each row is parsed independently, as Bubble Tea can redraw any row.
			for row, line := range lines {
				readColors(line, func(r rune, fg, bg string) {
					wantBG := profile.Color(string(Surface)).Sequence(true)
					if row == 0 || row == len(lines)-1 {
						wantBG = profile.Color(string(PanelColor)).Sequence(true)
					}
					if r == 'Y' {
						wantBG = profile.Color(string(Primary)).Sequence(true)
					}
					if bg != wantBG {
						t.Fatalf("row %d, %q: background %q, want %q", row, r, bg, wantBG)
					}
					if row > 0 && row < len(lines)-1 && r != ' ' {
						wantFG := profile.Color(string(Fixed)).Sequence(false)
						if r == 'X' {
							wantFG = profile.Color(string(Secondary)).Sequence(false)
						}
						if fg != wantFG {
							t.Fatalf("%q: foreground %q, want %q", r, fg, wantFG)
						}
					}
				})
			}
		})
	}
}

func TestScreenRespectsNoColor(t *testing.T) {
	previous := lipgloss.ColorProfile()
	t.Cleanup(func() { lipgloss.SetColorProfile(previous) })
	lipgloss.SetColorProfile(termenv.Ascii)
	if strings.Contains(Screen(40, 8, "TEST", CmdStyle.Render("Move"), "Quit"), "\x1b[") {
		t.Fatal("screen added colors in no-color mode")
	}
}

func TestSudokuMatchingValuesHaveReadableContrast(t *testing.T) {
	previous := lipgloss.ColorProfile()
	t.Cleanup(func() { lipgloss.SetColorProfile(previous) })
	lipgloss.SetColorProfile(termenv.TrueColor)
	for _, editable := range []bool{false, true} {
		cell := CellStyle(false, editable, false, true, false).Render("5")
		readColors(cell, func(r rune, fg, bg string) {
			if r == '5' && (fg != termenv.RGBColor(Surface).Sequence(false) || bg != termenv.RGBColor(Warning).Sequence(true)) {
				t.Fatalf("matching digit has poor contrast: fg %s bg %s", fg, bg)
			}
		})
	}
}
