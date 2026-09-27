package game2048

const size = 4

type Board [size][size]int

type direction int

const (
	left direction = iota
	right
	up
	down
)

// Merge toward the start of the line. A newly merged tile cannot merge again
// until the next move: [2, 2, 4, 0] becomes [4, 4, 0, 0].
func slideLine(line [size]int) (result [size]int, score int, merged [size]bool) {
	var tiles []int
	for _, value := range line {
		if value != 0 {
			tiles = append(tiles, value)
		}
	}
	for source, target := 0, 0; source < len(tiles); target++ {
		value := tiles[source]
		if source+1 < len(tiles) && value == tiles[source+1] {
			value *= 2
			score += value
			merged[target] = true
			source++
		}
		result[target] = value
		source++
	}
	return
}

func coordinates(d direction, line, position int) (row, column int) {
	switch d {
	case right:
		return line, size - 1 - position
	case up:
		return position, line
	case down:
		return size - 1 - position, line
	default:
		return line, position
	}
}

func slide(board Board, d direction) (result Board, score int, merged [size][size]bool) {
	for line := 0; line < size; line++ {
		var values [size]int
		for position := 0; position < size; position++ {
			row, column := coordinates(d, line, position)
			values[position] = board[row][column]
		}
		values, gain, marks := slideLine(values)
		score += gain
		for position := 0; position < size; position++ {
			row, column := coordinates(d, line, position)
			result[row][column] = values[position]
			merged[row][column] = marks[position]
		}
	}
	return
}

func (b Board) canMove() bool {
	for row := range b {
		for column, value := range b[row] {
			if value == 0 || (column+1 < size && value == b[row][column+1]) ||
				(row+1 < size && value == b[row+1][column]) {
				return true
			}
		}
	}
	return false
}

func (b Board) highest() int {
	value := 0
	for _, row := range b {
		for _, tile := range row {
			value = max(value, tile)
		}
	}
	return value
}
