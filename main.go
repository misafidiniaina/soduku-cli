package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
	"github.com/misafidiniaina/sudoku/internal/ui"
)


type Model struct {

	// Data for the game 
	Cells [9][9]int
	

	// Postion of the cursor index 0 represting the x axes and index 1 for the y axes
	cursor [2]int
}


func (m Model) Init() tea.Cmd {
	return  tea.ClearScreen
}






func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	i := m.cursor[0]
	j := m.cursor[1]

	switch msg := msg.(type) {
	case tea.KeyMsg:

		key := msg.String()

		switch key{

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up":
			m.cursor[1] = logic.CursorHandling("up", m.cursor[1])

		case "down":
			m.cursor[1] = logic.CursorHandling("down", m.cursor[1])
		
		case "left":
			m.cursor[0] = logic.CursorHandling("left", m.cursor[0])

		case "right":
			m.cursor[0] = logic.CursorHandling("right", m.cursor[0])

		case "backspace":
			m.Cells[j][i] = 0
		
		case "delete":
			m.Cells[j][i] = 0
		

		default:
			if len(key) == 1 && key[0] >= '1' && key[0] <= '9' && logic.IsEditable(m.cursor){
				m.Cells[j][i] = int(key[0] - '0')
				//because the key is an ASCII
			}
			
		}

	}

	return m, nil
}


func (m Model) View() string {

	GameView := lipgloss.JoinVertical(
		lipgloss.Left,
		ui.GameHeader(125, 2, "12:23"),
		ui.GameBoard(m.Cells, m.cursor),
		ui.CommandHelper(),
	)

	wrapper := ui.WrapperStyle.Render(GameView) + "\n"
	return wrapper
}



func main() {

	data := logic.GenerateData()
	p := tea.NewProgram(Model{Cells: data, cursor: [2]int{0,0}},)

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}