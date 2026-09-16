package tetris

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	boardWidth  = 10
	boardHeight = 20
	baseTick    = 700 * time.Millisecond
	minimumTick = 100 * time.Millisecond
)

type Point struct {
	X int
	Y int
}

type PieceType int

const (
	I PieceType = iota
	O
	T
	J
	L
	S
	Z
)

type Piece struct {
	Type     PieceType
	Rotation int
	X        int
	Y        int
}

type tickMsg time.Time

type Model struct {
	Board    [boardHeight][boardWidth]PieceType
	Current  Piece
	Next     Piece
	Score    int
	Lines    int
	Level    int
	GameOver bool
	Paused   bool

	random *rand.Rand
}

var shapes = [7][4]Point{
	{{0, 1}, {1, 1}, {2, 1}, {3, 1}},
	{{1, 0}, {2, 0}, {1, 1}, {2, 1}},
	{{1, 0}, {0, 1}, {1, 1}, {2, 1}},
	{{0, 0}, {0, 1}, {1, 1}, {2, 1}},
	{{2, 0}, {0, 1}, {1, 1}, {2, 1}},
	{{1, 0}, {2, 0}, {0, 1}, {1, 1}},
	{{0, 0}, {1, 0}, {1, 1}, {2, 1}},
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
	m.Board = [boardHeight][boardWidth]PieceType{}
	m.Score = 0
	m.Lines = 0
	m.Level = 1
	m.GameOver = false
	m.Paused = false
	m.Current = m.newPiece()
	m.Next = m.newPiece()
}

func (m Model) Init() tea.Cmd {
	return tick(m.Level)
}

func tick(level int) tea.Cmd {
	return tea.Tick(tickRate(level), func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func tickRate(level int) time.Duration {
	rate := baseTick - time.Duration(level-1)*60*time.Millisecond
	if rate < minimumTick {
		return minimumTick
	}
	return rate
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tickMsg:
		if !m.Paused && !m.GameOver {
			m.dropOne()
		}
		return m, tick(m.Level)

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
		case "left", "a":
			m.move(-1, 0)
		case "right", "d":
			m.move(1, 0)
		case "down", "s":
			if m.dropOne() {
				m.Score++
			}
		case "up", "w":
			m.rotate()
		case "enter":
			m.hardDrop()
		}
	}

	return m, nil
}

func (m Model) newPiece() Piece {
	return Piece{Type: PieceType(m.random.Intn(len(shapes))), X: 3}
}

func (m Model) cells(piece Piece) []Point {
	cells := make([]Point, len(shapes[piece.Type]))
	for index, point := range shapes[piece.Type] {
		x, y := point.X, point.Y
		if piece.Type != O {
			for rotation := 0; rotation < piece.Rotation%4; rotation++ {
				x, y = 3-y, x
			}
		}
		cells[index] = Point{X: piece.X + x, Y: piece.Y + y}
	}
	return cells
}

func (m Model) valid(piece Piece) bool {
	for _, point := range m.cells(piece) {
		if point.X < 0 || point.X >= boardWidth || point.Y >= boardHeight {
			return false
		}
		if point.Y >= 0 && m.Board[point.Y][point.X] != 0 {
			return false
		}
	}
	return true
}

func (m *Model) move(dx, dy int) bool {
	if m.GameOver || m.Paused {
		return false
	}
	candidate := m.Current
	candidate.X += dx
	candidate.Y += dy
	if !m.valid(candidate) {
		return false
	}
	m.Current = candidate
	return true
}

func (m *Model) rotate() bool {
	candidate := m.Current
	candidate.Rotation++
	if m.valid(candidate) {
		m.Current = candidate
		return true
	}
	return false
}

func (m *Model) dropOne() bool {
	if m.move(0, 1) {
		return true
	}
	if m.GameOver || m.Paused {
		return false
	}
	m.lock()
	return false
}

func (m *Model) hardDrop() {
	if m.GameOver || m.Paused {
		return
	}
	distance := 0
	for m.move(0, 1) {
		distance++
	}
	m.Score += distance * 2
	m.lock()
}

func (m *Model) lock() {
	for _, point := range m.cells(m.Current) {
		if point.Y < 0 {
			m.GameOver = true
			return
		}
		m.Board[point.Y][point.X] = m.Current.Type + 1
	}

	m.clearLines()
	m.Current = m.Next
	m.Current.X = 3
	m.Current.Y = 0
	m.Current.Rotation = 0
	m.Next = m.newPiece()
	if !m.valid(m.Current) {
		m.GameOver = true
	}
}

func (m *Model) clearLines() {
	cleared := 0
	for row := boardHeight - 1; row >= 0; row-- {
		full := true
		for column := 0; column < boardWidth; column++ {
			if m.Board[row][column] == 0 {
				full = false
				break
			}
		}
		if !full {
			continue
		}
		cleared++
		for current := row; current > 0; current-- {
			m.Board[current] = m.Board[current-1]
		}
		m.Board[0] = [boardWidth]PieceType{}
		row++
	}

	if cleared == 0 {
		return
	}
	m.Lines += cleared
	m.Level = 1 + m.Lines/10
	lineScores := [...]int{0, 100, 300, 500, 800}
	m.Score += lineScores[cleared] * m.Level
}

func (m Model) View() string {
	var rows []string
	for row := 0; row < boardHeight; row++ {
		var line strings.Builder
		line.WriteString("│")
		for column := 0; column < boardWidth; column++ {
			value := m.Board[row][column]
			for _, point := range m.cells(m.Current) {
				if point.X == column && point.Y == row {
					value = m.Current.Type + 1
				}
			}
			line.WriteString(renderCell(value))
		}
		line.WriteString("│")
		rows = append(rows, line.String())
	}

	status := fmt.Sprintf("Score: %d   Lines: %d   Level: %d   [Arrows/WASD] Move/Rotate   [Enter] Drop   [P] Pause   [R] Restart   [Q] Quit", m.Score, m.Lines, m.Level)
	if m.GameOver {
		status = fmt.Sprintf("GAME OVER - Score: %d   Lines: %d   [R] Restart   [Q] Quit", m.Score, m.Lines)
	} else if m.Paused {
		status = "PAUSED   [P/Space] Resume   [R] Restart   [Q] Quit"
	}

	preview := lipgloss.NewStyle().Foreground(lipgloss.Color("#a1a1a1")).Render(fmt.Sprintf("Next: %s", pieceName(m.Next.Type)))
	board := lipgloss.JoinVertical(lipgloss.Left, rows...)
	return lipgloss.JoinVertical(lipgloss.Left, titleStyle.Render("TETRIS"), preview, board, statusStyle.Render(status))
}

func renderCell(value PieceType) string {
	if value == 0 {
		return "  "
	}
	colors := [...]lipgloss.Color{"#00d9ff", "#ffd166", "#c77dff", "#4d96ff", "#ff9f1c", "#06d6a0", "#ef476f"}
	return lipgloss.NewStyle().Foreground(colors[value-1]).Render("[]")
}

func pieceName(piece PieceType) string {
	return [...]string{"I", "O", "T", "J", "L", "S", "Z"}[piece]
}

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#ff9f1c"))
	statusStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#87CEEB"))
)
