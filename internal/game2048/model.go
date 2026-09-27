package game2048

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Board       Board
	Score       int
	BestScore   int
	Moves       int
	Width       int
	Paused      bool
	GameOver    bool
	Won         bool
	KeepPlaying bool

	random    *rand.Rand
	lastGain  int
	merged    [size][size]bool
	spawned   [2]int
	flashing  bool
	animation uint64
}

type settleMsg uint64

func NewModel() Model {
	return newModel(rand.New(rand.NewSource(time.Now().UnixNano())))
}

func newModel(random *rand.Rand) Model {
	m := Model{random: random, Width: 80}
	m.reset()
	return m
}

func (m *Model) reset() {
	m.Board = Board{}
	m.Score, m.Moves, m.lastGain = 0, 0, 0
	m.Paused, m.GameOver, m.Won, m.KeepPlaying = false, false, false, false
	m.flashing = false
	m.merged = [size][size]bool{}
	m.animation++
	m.spawn()
	m.spawn()
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = message.Width
	case settleMsg:
		if uint64(message) == m.animation {
			m.flashing = false
		}
	case tea.KeyMsg:
		key := message.String()
		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.reset()
			return m, nil
		case "p", " ", "space":
			if !m.GameOver {
				m.Paused = !m.Paused
			}
			return m, nil
		}
		if m.Paused || m.GameOver {
			return m, nil
		}
		if m.Won && !m.KeepPlaying {
			if key == "enter" {
				m.KeepPlaying = true
			}
			return m, nil
		}
		var d direction
		switch key {
		case "left", "a":
			d = left
		case "right", "d":
			d = right
		case "up", "w":
			d = up
		case "down", "s":
			d = down
		default:
			return m, nil
		}
		if m.move(d) {
			version := m.animation
			return m, tea.Tick(180*time.Millisecond, func(time.Time) tea.Msg {
				return settleMsg(version)
			})
		}
	}
	return m, nil
}

func (m *Model) move(d direction) bool {
	next, gain, merged := slide(m.Board, d)
	if next == m.Board {
		m.GameOver = !m.Board.canMove()
		return false
	}
	m.Board, m.merged, m.lastGain = next, merged, gain
	m.Score += gain
	m.BestScore = max(m.BestScore, m.Score)
	m.Moves++
	m.spawn()
	m.GameOver = !m.Board.canMove()
	m.Won = m.Won || m.Board.highest() >= 2048
	m.flashing = true
	m.animation++
	return true
}

func (m *Model) spawn() {
	var empty [][2]int
	for row := range m.Board {
		for column, value := range m.Board[row] {
			if value == 0 {
				empty = append(empty, [2]int{row, column})
			}
		}
	}
	if len(empty) == 0 {
		return
	}
	m.spawned = empty[m.random.Intn(len(empty))]
	value := 2
	if m.random.Intn(10) == 0 {
		value = 4
	}
	m.Board[m.spawned[0]][m.spawned[1]] = value
}
