package snake

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/soduku-cli/internal/ui"
)

const (
	boardWidth      = 24
	boardHeight     = 14
	initialTickRate = 150 * time.Millisecond
	minimumTickRate = 65 * time.Millisecond
	scorePerLevel   = 5
)

type Point struct {
	X int
	Y int
}

type Direction Point

var (
	up    = Direction{X: 0, Y: -1}
	down  = Direction{X: 0, Y: 1}
	left  = Direction{X: -1, Y: 0}
	right = Direction{X: 1, Y: 0}
)

type tickMsg time.Time

type Model struct {
	Snake     []Point
	Food      Point
	Direction Direction
	Score     int
	BestScore int
	Width     int
	GameOver  bool
	Paused    bool
	Won       bool
	turned    bool

	random *rand.Rand
}

func NewModel() Model {
	return newModel(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func newModel(random *rand.Rand) Model {
	model := Model{random: random}
	model.reset()
	return model
}

func (m *Model) reset() {
	m.Snake = []Point{{X: boardWidth / 2, Y: boardHeight / 2}, {X: boardWidth/2 - 1, Y: boardHeight / 2}}
	m.Direction = right
	m.Score = 0
	m.GameOver = false
	m.Won, m.turned = false, false
	m.Paused = false
	m.Food = m.randomFood()
}

func (m Model) Init() tea.Cmd {
	return tick(m.Score)
}

func tick(score int) tea.Cmd {
	return tea.Tick(tickRate(score), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func tickRate(score int) time.Duration {
	rate := initialTickRate - time.Duration(score/scorePerLevel)*10*time.Millisecond
	if rate < minimumTickRate {
		return minimumTickRate
	}
	return rate
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = message.Width
	case tickMsg:
		if !m.Paused && !m.GameOver {
			m.step()
		}
		return m, tick(m.Score)

	case tea.KeyMsg:
		switch message.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.reset()
		case "p", "space":
			if !m.GameOver {
				m.Paused = !m.Paused
			}
		case "up", "w":
			m.changeDirection(up)
		case "down", "s":
			m.changeDirection(down)
		case "left", "a":
			m.changeDirection(left)
		case "right", "d":
			m.changeDirection(right)
		}
	}

	return m, nil
}

func (m *Model) changeDirection(direction Direction) {
	if m.Paused || m.GameOver || m.turned || direction == m.Direction {
		return
	}
	if direction.X == -m.Direction.X && direction.Y == -m.Direction.Y {
		return
	}
	m.Direction = direction
	m.turned = true
}

func (m *Model) step() {
	m.turned = false
	head := m.Snake[0]
	next := Point{X: head.X + m.Direction.X, Y: head.Y + m.Direction.Y}

	if next.X < 0 || next.X >= boardWidth || next.Y < 0 || next.Y >= boardHeight || m.occupies(next) {
		if next.X >= 0 && next.X < boardWidth && next.Y >= 0 && next.Y < boardHeight && next == m.Snake[len(m.Snake)-1] && next != m.Food {
			m.move(next)
			return
		}
		m.GameOver = true
		return
	}

	m.move(next)
	if next == m.Food {
		m.Score++
		if m.Score > m.BestScore {
			m.BestScore = m.Score
		}
		if len(m.Snake) == boardWidth*boardHeight {
			m.Won, m.GameOver = true, true
			return
		}
		m.Food = m.randomFood()
	}
}

func (m *Model) move(next Point) {
	m.Snake = append([]Point{next}, m.Snake...)
	if next != m.Food {
		m.Snake = m.Snake[:len(m.Snake)-1]
	}
}

func (m Model) occupies(point Point) bool {
	for _, segment := range m.Snake {
		if segment == point {
			return true
		}
	}
	return false
}

func (m Model) randomFood() Point {
	free := make([]Point, 0, boardWidth*boardHeight-len(m.Snake))
	for y := 0; y < boardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			p := Point{X: x, Y: y}
			if !m.occupies(p) {
				free = append(free, p)
			}
		}
	}
	if len(free) == 0 {
		return Point{X: -1, Y: -1}
	}
	return free[m.random.Intn(len(free))]
}

func (m Model) View() string {
	rows := []string{"╭" + strings.Repeat("─", boardWidth*2) + "╮"}
	for y := 0; y < boardHeight; y++ {
		var row strings.Builder
		row.WriteString("│")
		for x := 0; x < boardWidth; x++ {
			point := Point{X: x, Y: y}
			cell := "  "
			if point == m.Food {
				cell = foodStyle.Render("● ")
			} else if point == m.Snake[0] {
				glyph := map[Direction]string{up: "▲ ", down: "▼ ", left: "◀ ", right: "▶ "}[m.Direction]
				cell = headStyle.Render(glyph)
			} else if m.occupies(point) {
				cell = bodyStyle.Render("■ ")
			}
			row.WriteString(cell)
		}
		row.WriteString("│")
		rows = append(rows, row.String())
	}

	rows = append(rows, "╰"+strings.Repeat("─", boardWidth*2)+"╯")
	level := 1 + m.Score/scorePerLevel
	progress := strings.Repeat("━", m.Score%scorePerLevel) + strings.Repeat("─", scorePerLevel-m.Score%scorePerLevel)
	stats := ui.State(m.Paused, m.GameOver, m.Won) + "\n\n" +
		ui.Stat("SCORE", m.Score) + "\n\n" +
		ui.Stat("SESSION BEST", m.BestScore) + "\n\n" +
		ui.Stat("LEVEL", fmt.Sprintf("%d  %s", level, progress)) + "\n\n" +
		ui.CmdStyle.Render("↑↓←→ / WASD Move\nSpace Pause")
	if m.GameOver {
		stats += "\nR to play again"
	}
	board := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return ui.Playfield(m.Width, board, ui.Panel("SNAKE", stats, 26))

}

var (
	headStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
	bodyStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#5f7aff"))
	foodStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff005d"))
)
