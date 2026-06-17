package logic

import "math/rand/v2"

func GenerateData() [9][9]int{
	return [9][9]int{
		{0, 4, 9, 0, 3, 1, 0, 0, 0},
		{0, 3, 0, 0, 7, 0, 9, 8, 0},
		{8, 0, 0, 0, 4, 0, 0, 0, 0},

		{0, 0, 6, 1, 5, 7, 8, 0, 2},
		{0, 5, 0, 2, 8, 4, 1, 0, 6},
		{0, 1, 0, 0, 0, 0, 0, 0, 5},

		{0, 8, 5, 7, 0, 3, 4, 0, 9},
		{0, 0, 2, 0, 0, 5, 3, 7, 0},
		{0, 0, 4, 0, 2, 0, 5, 0, 1},
	}
}

func SolvedDataGen() [9][9]int{
	row := [9]int{5,9,3,8,1,6,7,4,2}
	var result [9][9]int

	rand.Shuffle(len(row), func(i, j int) {
		row[i], row[j] = row[j], row[i]
	})


	// Get a basic valid soduku (shifting per 3 and add 1 shift in line is in 3/6th)
	rowtmp := row
	for i := range 9 {
		decalage := 0
		if (i+1)%3 == 0 {
			decalage = 1
		}else{
			decalage = 0
		}

		result[i] = rowtmp
		memory := rowtmp
		for j := range 9 {
			if(j>(5 -decalage)){
				rowtmp[j] = memory[j-6+decalage]
			}else{
				rowtmp[j] = memory[j+3+decalage]
			}
		}
	}



	// shifle group of 3 row
	firstgroupPerm := rand.Perm(3)
	reslutMemory := result
	for x := range 3 {
		result[x] = reslutMemory[firstgroupPerm[x]]
	}

	secondGroupPerm := rand.Perm(3)
	for x := 3; x < 6; x++ {
		result[x] = reslutMemory[secondGroupPerm[x-3]+3]
	}

	thirdGroupPerm := rand.Perm(3)
	for x := 6; x < 9; x++ {
		result[x] = reslutMemory[thirdGroupPerm[x-6]+6]
	}


	// shifle the per 3 column group
	ColPrem := rand.Perm(3)
	for i := range 9 {
		lineRow := result[i]
		for x := range 3 {
			result[i][x] = lineRow[(ColPrem[0]*3)+x]
		}
		for x := 3; x < 6; x++ {
			result[i][x] = lineRow[(ColPrem[1]*3)+(x-3)]
		}
		for x := 6; x < 9; x++ {
			result[i][x] = lineRow[(ColPrem[2]*3)+(x-6)]
		}
	}

	// shifle group of 3 column
	firstLineGroupPerm := rand.Perm(3)
	SecondLineGroupPerm := rand.Perm(3)
	ThirdLineGroupPerm := rand.Perm(3)
	for i := range 9 {
		lineRow := result[i]
		for x := range 3 {
			result[i][x] = lineRow[firstLineGroupPerm[x]]
		}
		for x := 3; x < 6; x++ {
			result[i][x] = lineRow[SecondLineGroupPerm[x-3]+3]
		}
		for x := 6; x < 9; x++ {
			result[i][x] = lineRow[ThirdLineGroupPerm[x-6]+6]
		}
	}



	// shifle the  per 3 row gourps
	LinePrem := rand.Perm(3)
	reslutMemory = result
	for x := range 3 {
		result[x] = reslutMemory[(LinePrem[0]*3)+x]
	}
	for x := 3; x < 6; x++ {
		result[x] = reslutMemory[(LinePrem[1]*3)+(x-3)]
	}
	for x := 6; x < 9; x++ {
		result[x] = reslutMemory[(LinePrem[2]*3)+(x-6)]
	}
	return  result
}

var(
	SolvedData = [9][9]int{
		{6, 4, 9, 8, 3, 1, 2, 5, 7},
		{5, 3, 1, 6, 7, 2, 9, 8, 4},
		{8, 2, 7, 5, 4, 9, 6, 1, 3},

		{4, 9, 6, 1, 5, 7, 8, 3, 2},
		{7, 5, 3, 2, 8, 4, 1, 9, 6},
		{2, 1, 8, 3, 9, 6, 7, 4, 5},

		{1, 8, 5, 7, 6, 3, 4, 2, 9},
		{9, 6, 2, 4, 1, 5, 3, 7, 8},
		{3, 7, 4, 9, 2, 8, 5, 6, 1},
	}
)

func IsNotValid(i int, j int, value int) bool {
	
	if value != SolvedData[i][j] {
		return true
	}
	return false
}

