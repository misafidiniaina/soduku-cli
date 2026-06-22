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


func isSodukuValid(sodukuData [9]int) bool{
	result := true
	for i := 1; i < 10; i++ {
		occurenceCounter := 0
		for j := 0; j < 9; j++ {
			if sodukuData[j] == i {
				occurenceCounter++
			}
			if occurenceCounter == 2 {
				result = false
				break
			}
		}
	}

	return  result
}