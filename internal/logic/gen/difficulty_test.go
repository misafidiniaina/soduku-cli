package gen

import "testing"

func TestDifficultyEmptyCells(t *testing.T) {
	want := map[Difficulty]int{Easy: 30, Medium: 40, Hard: 45, Expert: 50, Master: 55, Extreme: 60}
	for difficulty, cells := range want {
		if got := difficulty.EmptyCells(); got != cells {
			t.Errorf("%s has %d empty cells, want %d", difficulty, got, cells)
		}
	}
}

func TestParseDifficulty(t *testing.T) {
	for _, value := range []string{"easy", "medium", "hard", "expert", "master", "extreme"} {
		if _, err := ParseDifficulty(value); err != nil {
			t.Errorf("ParseDifficulty(%q) returned error: %v", value, err)
		}
	}
	if _, err := ParseDifficulty("impossible"); err == nil {
		t.Fatal("ParseDifficulty accepted an invalid level")
	}
}
