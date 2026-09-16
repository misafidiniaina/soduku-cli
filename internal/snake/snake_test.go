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
