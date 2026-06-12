package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


type Model struct {
	Cells [9][9]int
	cursorPostion [2]int
}


func (m Model) Init() tea.Cmd {
	return  nil
}




func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String(){

		case "q", "ctrl+c":
			return m, tea.Quit

		case "up":
			
			m.Cells[8][8] = 2
			
		}

	}

	return m, nil
}


func (m Model) View() string {
	var gameTable string

	gameTable = "\nWelcome to our sodduku game and have fun\n\n"

	counter := 1

	for {
		row := 1
		for {
			if row <= 9 {

				if row == 3 || row == 6{
					gameTable = gameTable + "-----   "
				} else{
					gameTable = gameTable + "----- "
				// gameTable = gameTable + fmt.Sprint(row) + fmt.Sprint(counter) + "  "
				}
			}else{
				break
			}
			row = row + 1
		}
		gameTable = gameTable + "\n"
		row = 1
		for {
			if row <= 9 {
				if row == 3 || row == 6{
					gameTable = gameTable +"| "+fmt.Sprint(m.Cells[counter-1][row-1])+" |   "
				} else{
					gameTable = gameTable +"| "+fmt.Sprint(m.Cells[counter-1][row-1])+" | "
				}
				// we will add the model value based on row and counter value
			}else{
				break
			}
			row = row + 1
		}
		
		
		if counter < 9 {
			if counter == 3 || counter == 6 || counter == 9{
				gameTable = gameTable + "\n----- ----- -----   ----- ----- -----   ----- ----- ----- \n\n"
				
			} else{
				gameTable = gameTable + "\n"
			}
			
		}else {
			break
		}
		counter = counter + 1
	}
	gameTable = gameTable + "\n----- ----- -----   ----- ----- -----   ----- ----- ----- \n\n"
	return fmt.Sprintf("%v",gameTable)
}



func main() {

	
	p := tea.NewProgram(Model{Cells: [9][9]int{
			{0,4,9, 0,3,0 ,8,0,0},{0,3,1, 0,7,0 ,0,4,0},{0,0,0, 9,8,0, 0,0,0},
			{0,0,6, 0,5,0, 0,1,0},{1,5,7, 2,8,4, 0,0,0},{8,0,2 ,1,0,6, 0,0,5},
			{0,8,5, 0,0,2, 0,0,4},{7,0,3 ,0,0,5, 0,2,0},{4,0,9 ,3,7,0, 5,0,1},
			// create a function that will generate such data
		},
		cursorPostion: [2]int{0,0},
		
	})

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}