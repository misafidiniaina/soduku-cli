package snake

import (
	"math/rand"
	"testing"
)

func TestStepMovesSnake(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	initialLength := len(model.Snake)
	initialHead := model.Snake[0]

	model.step()

	if model.Snake[0] != (Point{X: initialHead.X + 1, Y: initialHead.Y}) {
		t.Fatalf("head = %+v, want one cell to the right", model.Snake[0])
	}
	if len(model.Snake) != initialLength {
		t.Fatalf("snake length = %d, want %d", len(model.Snake), initialLength)
	}
}

func TestStepGrowsWhenFoodIsAhead(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	model.Food = Point{X: model.Snake[0].X + 1, Y: model.Snake[0].Y}

	model.step()

	if len(model.Snake) != 3 {
		t.Fatalf("snake length = %d, want 3", len(model.Snake))
	}
	if model.Score != 1 {
		t.Fatalf("score = %d, want 1", model.Score)
	}
	if model.BestScore != 1 {
		t.Fatalf("best score = %d, want 1", model.BestScore)
	}
}

func TestStepEndsAtWall(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	model.Snake[0] = Point{X: 0, Y: model.Snake[0].Y}
	model.Direction = left

	model.step()

	if !model.GameOver {
		t.Fatal("GameOver = false, want true")
	}
}

func TestChangeDirectionRejectsReverse(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))

	model.changeDirection(left)

	if model.Direction != right {
		t.Fatalf("direction = %+v, want right", model.Direction)
	}
}

func TestStepAllowsMovingIntoTail(t *testing.T) {
	model := newModel(rand.New(rand.NewSource(1)))
	model.Snake = []Point{{X: 5, Y: 5}, {X: 5, Y: 4}, {X: 4, Y: 4}, {X: 4, Y: 5}}
	model.Direction = left
	model.Food = Point{X: 20, Y: 10}

	model.step()

	if model.GameOver {
		t.Fatal("moving into the tail should not end the game")
	}
}

func TestTickRateGetsFasterWithScore(t *testing.T) {
	if tickRate(0) <= tickRate(scorePerLevel) {
		t.Fatal("tick rate should decrease as the score increases")
	}
	if tickRate(100) != minimumTickRate {
		t.Fatalf("tick rate = %s, want minimum %s", tickRate(100), minimumTickRate)
	}
}
