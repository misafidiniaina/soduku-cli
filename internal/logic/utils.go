package logic

func CursorHandling(Move string, position int) int{
	var result int
	if Move == "left" {
		if position == 0 {
			result = 8
		}else{
			result = position - 1
		}
		
	}

	if Move == "down" {
		if position == 8 {
			result = 0
		}else{
			result = position + 1
		}
		
	}

	if Move == "up" {
		if position == 0 {
			result = 8
		}else{
			result = position - 1
		}
		
	}

	if Move == "right" {
		if position == 8 {
			result = 0
		}else{
			result = position + 1
		}
		
	}
	return result
}


func IsEditable(postion [2]int) bool{
	return GenerateData()[postion[1]][postion[0]] == 0
}

func GetValueInCursor(Data [9][9]int, cursor [2]int) int{
	return Data[cursor[1]][cursor[0]]
}