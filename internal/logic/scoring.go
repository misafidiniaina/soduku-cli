package logic

import "time"

// Score rewards correct entries, with penalties for mistakes and elapsed time.
func Score(cells, puzzle [9][9]int, mistakes int, elapsed time.Duration) int {
	correct := 0
	for row := range 9 {
		for column := range 9 {
			if puzzle[row][column] == 0 && cells[row][column] > 0 {
				correct++
			}
		}
	}

	score := correct*100 - mistakes*50 - int(elapsed.Seconds())
	if score < 0 {
		return 0
	}
	return score
}
