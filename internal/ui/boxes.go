package ui


func GameBoard(Data [9][9]int) string{
	var result string
	for i := 0; i < 9; i++ {
		result = result + Line(Data[i]) +"\n"
	}
	return result
}

