package minesweeper

import "math/rand"

const (
	maxBoardWidth  = 30
	maxBoardHeight = 16
)

type Point struct {
	Row int
	Col int
}

type Cell struct {
	Mine     bool
	Adjacent int
	Revealed bool
	Flagged  bool
}

type Board struct {
	Cells     [maxBoardHeight][maxBoardWidth]Cell
	Width     int
	Height    int
	Mines     int
	Flags     int
	Revealed  int
	Generated bool
}

func NewBoard(width, height, mines int) Board {
	return Board{Width: width, Height: height, Mines: mines}
}

func (b *Board) Reveal(point Point, random *rand.Rand) (hitMine, changed bool) {
	if !b.inBounds(point) {
		return false, false
	}
	cell := &b.Cells[point.Row][point.Col]
	if cell.Revealed || cell.Flagged {
		return false, false
	}
	if !b.Generated {
		b.placeMines(point, random)
	}
	cell = &b.Cells[point.Row][point.Col]
	if cell.Mine {
		cell.Revealed = true
		return true, true
	}
	b.revealSafeRegion(point)
	return false, true
}

func (b *Board) ToggleFlag(point Point) bool {
	if !b.inBounds(point) {
		return false
	}
	cell := &b.Cells[point.Row][point.Col]
	if cell.Revealed {
		return false
	}
	if cell.Flagged {
		cell.Flagged = false
		b.Flags--
		return true
	}
	if b.Flags >= b.Mines {
		return false
	}
	cell.Flagged = true
	b.Flags++
	return true
}

func (b Board) Won() bool {
	return b.Generated && b.Revealed == b.Width*b.Height-b.Mines
}

func (b Board) inBounds(point Point) bool {
	return point.Row >= 0 && point.Row < b.Height && point.Col >= 0 && point.Col < b.Width
}

func (b *Board) placeMines(first Point, random *rand.Rand) {
	positions := make([]Point, 0, b.Width*b.Height-1)
	for row := 0; row < b.Height; row++ {
		for col := 0; col < b.Width; col++ {
			point := Point{Row: row, Col: col}
			if point != first {
				positions = append(positions, point)
			}
		}
	}
	permutation := random.Perm(len(positions))
	for _, index := range permutation[:b.Mines] {
		point := positions[index]
		b.Cells[point.Row][point.Col].Mine = true
	}
	b.Generated = true

	for row := 0; row < b.Height; row++ {
		for col := 0; col < b.Width; col++ {
			if b.Cells[row][col].Mine {
				continue
			}
			count := 0
			for _, neighbor := range b.neighbors(Point{Row: row, Col: col}) {
				if b.Cells[neighbor.Row][neighbor.Col].Mine {
					count++
				}
			}
			b.Cells[row][col].Adjacent = count
		}
	}
}

func (b *Board) revealSafeRegion(start Point) {
	queue := []Point{start}
	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]
		cell := &b.Cells[point.Row][point.Col]
		if cell.Revealed || cell.Flagged || cell.Mine {
			continue
		}
		cell.Revealed = true
		b.Revealed++
		if cell.Adjacent == 0 {
			queue = append(queue, b.neighbors(point)...)
		}
	}
}

func (b Board) neighbors(point Point) []Point {
	neighbors := make([]Point, 0, 8)
	for row := max(0, point.Row-1); row <= min(b.Height-1, point.Row+1); row++ {
		for col := max(0, point.Col-1); col <= min(b.Width-1, point.Col+1); col++ {
			if row != point.Row || col != point.Col {
				neighbors = append(neighbors, Point{Row: row, Col: col})
			}
		}
	}
	return neighbors
}
