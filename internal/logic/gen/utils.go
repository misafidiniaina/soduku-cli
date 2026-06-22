package gen

import "slices"

func MissingValue(group [9]int) []int{
	var result []int
	normal := []int{1,2,3,4,5,6,7,8,9}

	var PresentValues []int
	for i := range len(group) {
		if group[i] != 0 {
			PresentValues = append(PresentValues, group[i])
		}
	}

	// build result as values from normal that are not present
	for _, v := range normal {
		skip := slices.Contains(PresentValues, v)
		if !skip {
			result = append(result, v)
		}
	}

	return result
}


// utils fucntion to know 	the case where the cell is
func WhichTranch(x int, y int) (tranchx int, tranchy int){
	var valuex int
	var valuey int
	
	if x < 3 {
		valuex = 0
	}else if x >= 3 && x < 6{
		valuex = 3
	}else {
		valuex = 6
	}

	if y < 3 {
		valuey = 0
	}else if y >= 3 && y < 6{
		valuey = 3
	}else {
		valuey = 6
	}

	return valuex,valuey
}


func IsArrayValid(arrayData [9]int) bool {
	seen := make(map[int]bool)

	for _, value := range arrayData {
		if value == 0 {
			continue
		}

		if seen[value] {
			return false
		}

		seen[value] = true
	}

	return true
}

func IsSodukuValid(sodukuData [9][9]int) bool{
	

	for i := 0; i < 9; i++ {
		if !IsArrayValid(sodukuData[i]) {
			return false
		}
	}

	// Check columns
	for i := 0; i < 9; i++ {
		var column [9]int

		for j := 0; j < 9; j++ {
			column[j] = sodukuData[j][i]
		}

		if !IsArrayValid(column) {
			return false
		}
	}

	// check 3x3 boxes validation
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			var box [9]int
			index := 0

			for i := boxRow * 3; i < boxRow*3+3; i++ {
				for j := boxCol * 3; j < boxCol*3+3; j++ {
					box[index] = sodukuData[i][j]
					index++
				}
			}

			if !IsArrayValid(box) {
				return false
			}
		}
	}


	return true
}