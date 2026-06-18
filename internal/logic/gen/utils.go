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