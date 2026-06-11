package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)


type Model struct {
	Cells [9][9]int
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
					gameTable = gameTable +"| "+fmt.Sprint(counter)+" |   "
				} else{
					gameTable = gameTable +"| "+fmt.Sprint(row)+" | "
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
	p := tea.NewProgram(Model{})

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}