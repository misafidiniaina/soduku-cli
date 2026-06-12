package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/misafidiniaina/sudoku/internal/logic"
)


type Model struct {

	// Data for the game 
	Cells [9][9]int

	// Postion of the cursor index 0 represting the x axes and index 1 for the y axes
	cursor [2]int
}


func (m Model) Init() tea.Cmd {
	return  nil
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
			if len(key) == 1 && key[0] >= '1' && key[0] <= '9' {
				m.Cells[j][i] = int(key[0] - '0')
				//because the key is an ASCII
			}
			
		}

	}

	return m, nil
}


func (m Model) View() string {
	// var gameTable string

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#ff0000")).
		Border(lipgloss.ThickBorder())

	footerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00ff00"))


	bigblockStyle := lipgloss.NewStyle().
		Width(60).
		Border(lipgloss.NormalBorder()).
		BorderBackground(lipgloss.Color("#0dfsdf")).
		Background(lipgloss.Color("62")).
		Foreground(lipgloss.Color("#000")).
		Bold(true)

	

	header := headerStyle.Render("salut mes loulous")
	footer := footerStyle.Render("this is the footer")
	

	result := lipgloss.JoinHorizontal(
		lipgloss.Center,
		footer,
		header,
	)

	finalresult := lipgloss.JoinVertical(
		lipgloss.Left,
		result,
		bigblockStyle.Render("salut"),
		"\n\n",
	)

	



	// gameTable = "\nWelcome to our sodduku game and have fun\n\n"


	// counter := 1

	// for {
	// 	row := 1
	// 	for {
	// 		if row <= 9 {

	// 			if row == 3 || row == 6{
	// 				gameTable = gameTable + "-----   "
	// 			} else{
	// 				gameTable = gameTable + "----- "
	// 			// gameTable = gameTable + fmt.Sprint(row) + fmt.Sprint(counter) + "  "
	// 			}
	// 		}else{
	// 			break
	// 		}
	// 		row = row + 1
	// 	}
	// 	gameTable = gameTable + "\n"
	// 	row = 1
	// 	for {
	// 		if row <= 9 {
	// 			var value string
	// 			// counter-1 and row-1 because the row and counter start with 1 not 0 like the data (m.Cells) index
	// 			if m.Cells[counter-1][row-1] == 0 {
	// 				value = " "
	// 			}else{
	// 				value = fmt.Sprint(m.Cells[counter-1][row-1])
	// 			}

	// 			if row == 3 || row == 6{
	// 				gameTable = gameTable +"| "+value+" |   "
	// 			} else{
	// 				gameTable = gameTable +"| "+value+" | "
	// 			}
	// 			// we will add the model value based on row and counter value
	// 		}else{
	// 			break
	// 		}
	// 		row = row + 1
	// 	}
		
		
	// 	if counter < 9 {
	// 		if counter == 3 || counter == 6 || counter == 9{
	// 			gameTable = gameTable + "\n----- ----- -----   ----- ----- -----   ----- ----- ----- \n\n"
				
	// 		} else{
	// 			gameTable = gameTable + "\n"
	// 		}
			
	// 	}else {
	// 		break
	// 	}
	// 	counter = counter + 1
	// }
	// gameTable = gameTable + "\n----- ----- -----   ----- ----- -----   ----- ----- ----- \n\n"
	// gameTable = gameTable + "The cursor position is in: " + fmt.Sprint(m.cursor)
	// return fmt.Sprintf("%v",gameTable)
	return finalresult
}



func main() {

	
	p := tea.NewProgram(Model{Cells: [9][9]int{
			{0,4,9, 0,3,0 ,8,0,0},
			{0,3,1, 0,7,0 ,0,4,0},
			{0,0,0, 9,8,0, 0,0,0},

			{0,0,6, 0,5,0, 0,1,0},
			{1,5,7, 2,8,4, 0,0,0},
			{8,0,2 ,1,0,6, 0,0,5},

			{0,8,5, 0,0,2, 0,0,4},
			{7,0,3 ,0,0,5, 0,2,0},
			{4,0,9 ,3,7,0, 5,0,1},
			// create a function that will generate such data
		},
		cursor: [2]int{0,0},
		
	})

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}