package game2048

import "testing"

func TestSlideLine(t *testing.T) {
	for _, test := range []struct {
		name  string
		input [4]int
		want  [4]int
		score int
	}{
		{"empty", [4]int{}, [4]int{}, 0},
		{"compress gaps", [4]int{0, 2, 0, 4}, [4]int{2, 4, 0, 0}, 0},
		{"merge across gaps", [4]int{2, 0, 2, 0}, [4]int{4, 0, 0, 0}, 4},
		{"no chain merge", [4]int{2, 2, 4, 0}, [4]int{4, 4, 0, 0}, 4},
		{"two pairs", [4]int{2, 2, 2, 2}, [4]int{4, 4, 0, 0}, 8},
		{"three equal tiles", [4]int{4, 4, 4, 0}, [4]int{8, 4, 0, 0}, 8},
		{"middle pair", [4]int{2, 4, 4, 2}, [4]int{2, 8, 2, 0}, 8},
		{"already packed", [4]int{2, 4, 8, 16}, [4]int{2, 4, 8, 16}, 0},
	} {
		t.Run(test.name, func(t *testing.T) {
			got, score, marks := slideLine(test.input)
			if got != test.want || score != test.score {
				t.Fatalf("got %v / %d, want %v / %d", got, score, test.want, test.score)
			}
			markedScore := 0
			for i, marked := range marks {
				if marked {
					markedScore += got[i]
				}
			}
			if markedScore != score {
				t.Fatal("merge highlights disagree with scored tiles")
			}
		})
	}
}

func TestSlideAllDirections(t *testing.T) {
	board := Board{{2, 2}, {2, 2}}
	for _, test := range []struct {
		d    direction
		want Board
	}{
		{left, Board{{4}, {4}}},
		{right, Board{{0, 0, 0, 4}, {0, 0, 0, 4}}},
		{up, Board{{4, 4}}},
		{down, Board{{}, {}, {}, {4, 4}}},
	} {
		got, score, _ := slide(board, test.d)
		if got != test.want || score != 8 {
			t.Fatalf("direction %d: got %v / %d, want %v / 8", test.d, got, score, test.want)
		}
	}
}

func TestAvailableMoves(t *testing.T) {
	full := Board{{2, 4, 2, 4}, {4, 2, 4, 2}, {2, 4, 2, 4}, {4, 2, 4, 2}}
	if full.canMove() {
		t.Fatal("checkerboard has no legal moves")
	}
	for _, cell := range [][2]int{{0, 1}, {1, 0}} {
		board := full
		board[cell[0]][cell[1]] = 2
		if !board.canMove() {
			t.Fatal("full board still has an adjacent pair")
		}
	}
	full[3][3] = 0
	if !full.canMove() {
		t.Fatal("empty cell should allow a move")
	}
}
