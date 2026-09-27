package sudoku

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku/logic"
	"github.com/misafidiniaina/soduku-cli/internal/sudoku/logic/gen"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
)

var difficulties = []gen.Difficulty{gen.Easy, gen.Medium, gen.Hard, gen.Expert, gen.Master, gen.Extreme}

type Model struct {
	Cells    [9][9]int
	Puzzle   [9][9]int
	Solution [9][9]int

	StartTime time.Time
	Elapsed   time.Duration
	Paused    bool

	Mistake        int
	GameOver       bool
	Won            bool
	Restarting     bool
	RestartChoice  int
	Width          int
	Height         int
	SelectingLevel bool
	LevelIndex     int
	cursor         [2]int
}

type tickMsg time.Time

func NewModel() Model {
	return Model{SelectingLevel: true, LevelIndex: 1, Width: 80, Height: 24}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, tick())
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	i := m.cursor[0]
	j := m.cursor[1]

	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = message.Width, message.Height

	case tickMsg:
		if !m.Paused && !m.SelectingLevel && !m.Restarting && !m.GameOver && !m.Won {
			m.Elapsed += time.Second
		}
		return m, tick()

	case tea.KeyMsg:
		key := message.String()
		if key == "q" || key == "ctrl+c" {
			return m, tea.Quit
		}
		if m.Paused && !m.Restarting && key != "p" && key != "r" {
			return m, nil
		}
		if m.Restarting {
			switch key {
			case "up", "left":
				m.RestartChoice = (m.RestartChoice + 2) % 3
			case "down", "right":
				m.RestartChoice = (m.RestartChoice + 1) % 3
			case "enter":
				m.Paused = false
				switch m.RestartChoice {
				case 0:
					m.Cells, m.Mistake, m.Elapsed, m.Restarting, m.GameOver, m.Won = m.Puzzle, 0, 0, false, false, false
				case 1:
					m.Puzzle, m.Solution = gen.PuzzleGenAt(difficulties[m.LevelIndex])
					m.Cells, m.Mistake, m.Elapsed, m.Restarting, m.GameOver, m.Won = m.Puzzle, 0, 0, false, false, false
				case 2:
					m.SelectingLevel, m.Restarting = true, false
				}
			case "esc":
				m.Restarting = false
			case "q", "ctrl+c":
				return m, tea.Quit
			}
			return m, nil
		}
		if m.GameOver || m.Won {
			if key == "r" {
				m.Restarting = true
				return m, nil
			}
			if key == "q" || key == "ctrl+c" {
				return m, tea.Quit
			}
			return m, nil
		}

		if m.SelectingLevel {
			switch key {
			case "up", "left":
				m.LevelIndex = (m.LevelIndex + len(difficulties) - 1) % len(difficulties)
			case "down", "right":
				m.LevelIndex = (m.LevelIndex + 1) % len(difficulties)
			case "1", "2", "3", "4", "5", "6":
				m.LevelIndex = int(key[0] - '1')
			case "enter":
				m.Puzzle, m.Solution = gen.PuzzleGenAt(difficulties[m.LevelIndex])
				m.Cells = m.Puzzle
				m.StartTime = time.Now()
				m.Elapsed, m.Mistake, m.Paused = 0, 0, false
				m.SelectingLevel, m.GameOver, m.Won = false, false, false
			}
			return m, nil
		}

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.Restarting = true
			m.RestartChoice = 0
			return m, nil
		case "up":
			m.cursor[1] = logic.CursorHandling("up", m.cursor[1])
		case "down":
			m.cursor[1] = logic.CursorHandling("down", m.cursor[1])
		case "left":
			m.cursor[0] = logic.CursorHandling("left", m.cursor[0])
		case "right":
			m.cursor[0] = logic.CursorHandling("right", m.cursor[0])
		case "backspace", "delete":
			if logic.IsEditableAt(m.Puzzle, m.cursor) {
				m.Cells[j][i] = 0
			}
		case "p":
			m.Paused = !m.Paused
		default:
			if len(key) == 1 && key[0] >= '1' && key[0] <= '9' && logic.IsEditableAt(m.Puzzle, m.cursor) {
				if m.Solution[j][i] != int(key[0]-'0') {
					m.Cells[j][i] = -int(key[0] - '0')
					m.Mistake++
					if m.Mistake >= 3 {
						m.GameOver = true
					}
				} else {
					m.Cells[j][i] = int(key[0] - '0')
					m.Won = gameWon(m.Cells, m.Puzzle, m.Solution)
				}
			}
		}
	}

	return m, nil
}

func gameWon(cells, puzzle, solution [9][9]int) bool {
	for row := range 9 {
		for column := range 9 {
			if puzzle[row][column] == 0 && cells[row][column] != solution[row][column] {
				return false
			}
		}
	}
	return true
}

func (m Model) View() string {
	if m.SelectingLevel {
		return ui.WrapperStyle.MarginTop(0).Render(ui.LevelSelector(difficulties, m.LevelIndex))
	}
	if m.Restarting {
		return ui.WrapperStyle.MarginTop(0).Render(ui.RestartSelector(m.RestartChoice))
	}
	board := ui.GameBoard(m.Cells, m.Puzzle, m.cursor, 0)
	stats := ui.State(m.Paused, m.GameOver, m.Won) + "\n\n" +
		ui.Stat("SCORE", logic.Score(m.Cells, m.Puzzle, m.Mistake, m.Elapsed)) + "\n\n" +
		ui.Stat("TIME", logic.Chrono(m.Elapsed)) + "\n\n" +
		ui.Stat("DIFFICULTY", string(difficulties[m.LevelIndex])) + "\n\n" +
		ui.Stat("MISTAKES", fmt.Sprintf("%d / 3", m.Mistake)) + "\n\n" +
		strings.ReplaceAll(m.cellHelp(), " · ", "\n") + "\n\n" +
		ui.CmdStyle.Render("↑↓←→ Move\n1–9 Enter number\nBackspace Clear")
	if m.GameOver || m.Won {
		stats += "\n\nR to play again"
	}
	return ui.Playfield(m.Width, board, ui.Panel("SUDOKU", stats, 20))
}

// Candidates use only the visible grid, so assistance never reveals the solution.
func (m Model) cellHelp() string {
	x, y := m.cursor[0], m.cursor[1]
	filled := 0
	for _, row := range m.Cells {
		for _, value := range row {
			if value > 0 {
				filled++
			}
		}
	}
	prefix := fmt.Sprintf("%d/81 filled · Row %d, column %d", filled, y+1, x+1)
	if m.Paused || m.GameOver || m.Won {
		return prefix
	}
	if m.Puzzle[y][x] != 0 {
		return prefix + " · Given clue"
	}
	if m.Cells[y][x] > 0 {
		return prefix
	}
	var candidates []string
	for n := 1; n <= 9; n++ {
		valid := true
		for i := 0; i < 9; i++ {
			if m.Cells[y][i] == n || m.Cells[i][x] == n || m.Cells[y/3*3+i/3][x/3*3+i%3] == n {
				valid = false
			}
		}
		if valid {
			candidates = append(candidates, fmt.Sprint(n))
		}
	}
	if len(candidates) == 0 {
		return prefix + " · No candidates"
	}
	return prefix + " · Candidates: " + strings.Join(candidates, " ")
}
