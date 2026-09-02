package gen

import "fmt"

type Difficulty string

const (
	Easy    Difficulty = "easy"
	Medium  Difficulty = "medium"
	Hard    Difficulty = "hard"
	Expert  Difficulty = "expert"
	Master  Difficulty = "master"
	Extreme Difficulty = "extreme"
)

func ParseDifficulty(value string) (Difficulty, error) {
	difficulty := Difficulty(value)
	switch difficulty {
	case Easy, Medium, Hard, Expert, Master, Extreme:
		return difficulty, nil
	default:
		return "", fmt.Errorf("invalid difficulty %q", value)
	}
}

func (d Difficulty) EmptyCells() int {
	switch d {
	case Easy:
		return 30
	case Medium:
		return 40
	case Hard:
		return 45
	case Expert:
		return 50
	case Master:
		return 55
	case Extreme:
		return 60
	default:
		return Medium.EmptyCells()
	}
}
