package minesweeper

import (
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type difficulty struct {
	Name   string
	Width  int
	Height int
	Mines  int
}

var difficulties = []difficulty{
	{Name: "Beginner", Width: 9, Height: 9, Mines: 10},
	{Name: "Intermediate", Width: 16, Height: 16, Mines: 40},
	{Name: "Expert", Width: 30, Height: 16, Mines: 99},
}

type Model struct {
	Board               Board
	Cursor              Point
	DifficultyIndex     int
	SelectingDifficulty bool
	Width               int
	Paused              bool
	GameOver            bool
	Won                 bool
	Elapsed             time.Duration

	random  *rand.Rand
	started bool
}

type tickMsg time.Time

func NewModel() Model {
	return Model{
		SelectingDifficulty: true,
		Width:               80,
		random:              rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(now time.Time) tea.Msg {
		return tickMsg(now)
	})
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.WindowSizeMsg:
		m.Width = message.Width
	case tickMsg:
		if m.started && !m.Paused && !m.GameOver && !m.Won {
			m.Elapsed += time.Second
		}
		return m, tick()
	case tea.KeyMsg:
		key := message.String()
		if key == "q" || key == "ctrl+c" {
			return m, tea.Quit
		}
		if m.SelectingDifficulty {
			m.updateDifficulty(key)
			return m, nil
		}
		if key == "r" {
			m.reset()
			return m, nil
		}
		if key == "p" {
			if !m.GameOver && !m.Won {
				m.Paused = !m.Paused
			}
			return m, nil
		}
		if m.Paused || m.GameOver || m.Won {
			return m, nil
		}
		switch key {
		case "up", "w":
			m.Cursor.Row = max(0, m.Cursor.Row-1)
		case "down", "s":
			m.Cursor.Row = min(m.Board.Height-1, m.Cursor.Row+1)
		case "left", "a":
			m.Cursor.Col = max(0, m.Cursor.Col-1)
		case "right", "d":
			m.Cursor.Col = min(m.Board.Width-1, m.Cursor.Col+1)
		case "enter", " ", "space":
			hitMine, changed := m.Board.Reveal(m.Cursor, m.random)
			if changed && !m.started {
				m.started = true
			}
			if hitMine {
				m.GameOver = true
				m.revealMines()
			} else if m.Board.Won() {
				m.Won = true
			}
		case "f":
			m.Board.ToggleFlag(m.Cursor)
		}
	}
	return m, nil
}

func (m *Model) updateDifficulty(key string) {
	switch key {
	case "up", "left":
		m.DifficultyIndex = (m.DifficultyIndex + len(difficulties) - 1) % len(difficulties)
	case "down", "right":
		m.DifficultyIndex = (m.DifficultyIndex + 1) % len(difficulties)
	case "1", "2", "3":
		m.DifficultyIndex = int(key[0] - '1')
	case "enter":
		m.reset()
		m.SelectingDifficulty = false
	}
}

func (m *Model) reset() {
	difficulty := difficulties[m.DifficultyIndex]
	m.Board = NewBoard(difficulty.Width, difficulty.Height, difficulty.Mines)
	m.Cursor = Point{}
	m.Elapsed = 0
	m.Paused = false
	m.GameOver = false
	m.Won = false
	m.started = false
}

func (m *Model) revealMines() {
	for row := 0; row < m.Board.Height; row++ {
		for col := 0; col < m.Board.Width; col++ {
			cell := &m.Board.Cells[row][col]
			if cell.Mine {
				cell.Revealed = true
			}
		}
	}
}
