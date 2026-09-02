package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
	"github.com/misafidiniaina/sudoku/internal/logic/gen"
	"github.com/misafidiniaina/sudoku/internal/ui"
)

var difficulties = []gen.Difficulty{gen.Easy, gen.Medium, gen.Hard, gen.Expert, gen.Master, gen.Extreme}

type Model struct {

	// Data for the game
	Cells    [9][9]int
	Puzzle   [9][9]int
	Solution [9][9]int

	StartTime time.Time
	Elapsed   time.Duration
	Paused    bool

	// player mistake counter
	Mistake        int
	GameOver       bool
	Won            bool
	Restarting     bool
	RestartChoice  int
	Width          int
	Height         int
	SelectingLevel bool
	LevelIndex     int
	// Postion of the cursor index 0 represting the x axes and index 1 for the y axes
	cursor [2]int
}

// a custom struct to represent time trick (a time based command that triggers event at regular intervals)
type TickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tea.ClearScreen,
		tick(),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	i := m.cursor[0]
	j := m.cursor[1]

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height

	case TickMsg:
		if !m.Paused {
			m.Elapsed += time.Second
		}
		return m, tick()

	case tea.KeyMsg:
		key := msg.String()
		if m.Restarting {
			switch key {
			case "up", "left":
				m.RestartChoice = (m.RestartChoice + 2) % 3
			case "down", "right":
				m.RestartChoice = (m.RestartChoice + 1) % 3
			case "enter":
				switch m.RestartChoice {
				case 0:
					m.Cells, m.Mistake, m.Elapsed, m.Restarting, m.GameOver, m.Won = m.Puzzle, 0, 0, false, false, false
				case 1:
					m.Puzzle, m.Solution = gen.PuzzleGenAt(difficulties[m.LevelIndex])
					m.Cells, m.Mistake, m.Elapsed, m.Restarting, m.GameOver, m.Won = m.Puzzle, 0, 0, false, false, false
				case 2:
					m.SelectingLevel, m.Restarting = true, false
				}
			case "q", "ctrl+c":
				return m, tea.Quit
			}
			return m, nil
		}
		if m.GameOver {
			if key == "r" {
				m.Restarting = true
				return m, nil
			}
			if key == "q" || key == "ctrl+c" {
				return m, tea.Quit
			}
			return m, nil
		}
		if m.Won {
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
			m.Cells[j][i] = 0

		case "p":
			m.Paused = !m.Paused
			return m, nil

		default:
			if len(key) == 1 && key[0] >= '1' && key[0] <= '9' && logic.IsEditableAt(m.Puzzle, m.cursor) {
				// if the user input is a mistake return a negatif value (easy to track)
				if m.Solution[j][i] != int(key[0]-'0') {
					m.Cells[j][i] = -int(key[0] - '0') // negative = mistake
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
		return ui.WrapperStyle.Render(ui.LevelSelector(difficulties, m.LevelIndex)) + "\n"
	}
	if m.Restarting {
		return ui.WrapperStyle.Render(ui.RestartSelector(m.RestartChoice)) + "\n"
	}
	var MaingContent string
	if m.GameOver {
		MaingContent = ui.GameOverStyle.Render("GAME OVER\nYou made 3 mistakes\nPress 'r' to restart or 'q' to quit")
	} else if m.Won {
		MaingContent = ui.WinStyle.Render("YOU WIN!\nPress 'r' to play again or 'q' to quit")
	} else if m.Paused {
		MaingContent = ui.PausedGameSyle.Render("        GAME PAUSED, \nPress 'p' to resume the game")
	} else {
		boardWidth := m.Width
		if boardWidth < 1 {
			boardWidth = 80
		}
		if boardWidth > 117 {
			boardWidth = 117
		}
		MaingContent = ui.GameBoard(m.Cells, m.Puzzle, m.cursor, boardWidth)
	}

	GameView := lipgloss.JoinVertical(
		lipgloss.Left,
		ui.GameHeader(logic.Score(m.Cells, m.Puzzle, m.Mistake, m.Elapsed), string(difficulties[m.LevelIndex]), m.Mistake, logic.Chrono(m.Elapsed)),
		MaingContent,
		ui.CommandHelper(),
	)

	wrapper := ui.WrapperStyle.Render(GameView) + "\n"
	return wrapper
}

func main() {
	p := tea.NewProgram(Model{SelectingLevel: true, LevelIndex: 1, Width: 80, Height: 24})

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
