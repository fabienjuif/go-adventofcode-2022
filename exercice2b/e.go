package exercice2b

import (
	"fmt"
	"strings"
)

func Run() {
	fmt.Println("exercice 2a")
	fmt.Println(process(input))
}

func process(text string) string {
	lines := strings.Split(text, "\n")
	totalScore := 0

	for _, line := range lines {
		opponents := NewShapeFromOpponent(line[0:1])
		needTo := NewNeedTo(line[2:3])
		ours := GetOurHand(opponents, needTo)

		score := ProcessScore(opponents, ours)
		totalScore += score
	}

	return fmt.Sprintf("%d", totalScore)
}

func GetOurHand(shape Shape, needTo NeedTo) Shape {
	switch needTo {
	case NeedToDraw:
		return shape
	case NeedToLoose:
		return shape.GetDefeated()
	case NeedToWin:
		return shape.GetDefeater()
	}

	panic(fmt.Sprintf("unknown needTo: %v", needTo))
}

type Shape struct {
	slug          string
	score         int
	opponentValue string
}

func (s Shape) String() string {
	return s.slug
}

func (s Shape) GetDefeater() Shape {
	switch s {
	case ShapeRock:
		return ShapePaper
	case ShapePaper:
		return ShapeScissors
	case ShapeScissors:
		return ShapeRock
	}
	panic(fmt.Sprintf("unknown shape: %v", s))
}

func (s Shape) GetDefeated() Shape {
	switch s {
	case ShapeRock:
		return ShapeScissors
	case ShapePaper:
		return ShapeRock
	case ShapeScissors:
		return ShapePaper
	}
	panic(fmt.Sprintf("unknown shape: %v", s))
}

var (
	ShapeRock     = Shape{"rock", 1, "A"}
	ShapePaper    = Shape{"paper", 2, "B"}
	ShapeScissors = Shape{"scissors", 3, "C"}
	Shapes        = []Shape{ShapeRock, ShapePaper, ShapeScissors}
)

func NewShapeFromOpponent(o string) Shape {
	for _, shape := range Shapes {
		if shape.opponentValue == o {
			return shape
		}
	}

	panic(fmt.Sprintf("unknown hand for opponent: %s", o))
}

type NeedTo struct {
	slug    string
	crypted string
}

func (n NeedTo) String() string {
	return n.slug
}

var (
	NeedToLoose = NeedTo{"loose", "X"}
	NeedToDraw  = NeedTo{"draw", "Y"}
	NeedToWin   = NeedTo{"win", "Z"}
	NeedToS     = []NeedTo{NeedToLoose, NeedToDraw, NeedToWin}
)

func NewNeedTo(o string) NeedTo {
	for _, needTo := range NeedToS {
		if needTo.crypted == o {
			return needTo
		}
	}

	panic(fmt.Sprintf("unknown crypted value: %s", o))
}

// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
// return -1 if left wins, 1 if right wins, 0 for a draw
func PickWinner(left, right Shape) int {
	if left == right {
		return 0
	}

	if left == ShapeRock {
		if right == ShapeScissors {
			return -1
		}
		return 1
	}

	if left == ShapeScissors {
		if right == ShapePaper {
			return -1
		}
		return 1
	}

	if left == ShapePaper {
		if right == ShapeRock {
			return -1
		}
		return 1
	}

	panic(fmt.Sprintf("Unknown left hand: %v", left))
}

func ProcessScore(left, right Shape) int {
	winner := PickWinner(left, right)
	if winner == 0 {
		return 3 + right.score
	}
	if winner == 1 {
		return 6 + right.score
	}
	return 0 + right.score
}

var input = `C Y
C Y
B Y
A Z
B Z
A X
A Y
A Y
A X
A Y
B Y
A Y
B Y
B Y
B Z
B Z
B Z
B Z
A Y
B Z
A Y
B X
B Y
B X
A X
A X
B Z
A X
A X
B Z
B Z
B Y
B Z
B Z
B Z
B Y
A X
A X
B Z
A X
B X
B X
C Y
B Z
C X
A X
A Y
B Y
A Y
B X
A X
B Y
B Z
B Z
B Y
B Z
C Z
B X
B X
B Z
B Z
B Z
B Z
A Y
B X
A X
C X
B Y
B Z
A Y
B Z
B Z
B Y
B Y
B Z
B Y
B Z
A X
B X
B X
A X
A X
B Z
B Z
B Z
B Z
B Y
B X
B Z
A X
A Y
B Z
A Y
B Z
B Y
B Z
A X
B Y
A Y
B Z
A Z
A Y
A Y
C Y
B Z
B X
A Z
B Z
A X
B Z
A Z
B Z
A X
B Y
A X
B Y
B Z
B X
B X
B Z
B Z
B Z
A X
B Z
A X
B X
B Z
A X
C Z
B Z
B Z
B Y
B Y
B Y
B Z
B Z
A Y
B Z
B Z
C Y
C Z
A X
B Z
B X
B Z
B Z
B Y
A X
B Z
B Y
A Z
B Z
A X
A X
B Y
A Y
B Z
B Z
B X
B Y
A X
A Y
B X
C Z
A Y
B Z
B Z
B Z
A Z
C Y
B Y
B Y
B Z
C Y
B Y
B X
B X
B Z
C Z
A X
B Z
B Z
B Z
B Z
B X
B X
A X
A Z
A Z
A X
C Y
B X
A X
A Y
A X
B X
A Y
B Z
B Z
A Y
A X
B Y
B Z
B Z
A X
A Z
B Z
B X
A X
B Z
B Y
A Y
A Z
B X
A Y
B Z
A Z
B Z
A Y
B Z
B X
B Y
A Y
B Z
B Z
A X
A X
B Y
B Z
A X
B Z
B Z
B Z
B Z
B Z
A Z
B Z
B Z
B X
A Y
C X
B Z
B Y
B Y
B Z
B Z
B Z
B Y
B X
B Y
C X
B Z
A Z
A Y
C X
A X
B X
A X
B X
A Y
B Z
A Y
A Y
B Z
B X
B Z
A Y
B Z
B X
C Z
C X
C Z
B Y
B X
B Z
B Z
B Y
B Z
B Y
B Z
A X
B X
B Z
A X
B Z
B Y
B Z
A X
B Y
C Y
A Z
B Z
C Z
A Y
B Z
A Y
A Z
B Y
A X
A X
B Z
B Z
B Z
A Y
B Z
A Z
B Z
A Y
A Z
B Y
C Z
B Z
A Z
B X
B Z
B Y
B Z
A Z
A Z
B Z
B X
C Z
B X
B Z
B Y
A X
B Z
A X
B X
B Z
B Z
A X
B X
C X
C X
B X
B Y
B Z
B X
B Y
B Y
B Z
A Z
B Z
C X
A Y
C X
B Z
A Y
B Z
B Y
B Z
B X
A X
B X
B Z
A Y
A Y
B Z
B Z
B Z
A Y
B Z
B X
C X
B Z
B Z
C Y
B Z
C Y
B Z
C Y
C X
B Z
C Y
A Y
A Y
C Z
B Z
B X
B Z
B X
C Z
B Z
A Z
B Z
C Y
B Z
A X
A Y
B Y
B Y
B Z
A Y
B X
B Z
B Z
A Z
B X
A Y
A Y
B Z
B Z
B Z
C X
A Z
B X
C Z
B Y
B Z
B X
A Y
B Z
A X
B X
B Y
A Y
B Z
B Z
B Y
A Z
B Z
A X
B Y
A Y
A X
A Y
A X
A X
B X
B Y
B X
B Y
B Z
B Y
B X
A X
B Z
A Z
A X
B Z
B Y
A X
B Y
A X
B X
C Y
B X
B X
B X
C Y
B X
B Y
B Y
B Y
B Y
B X
A Y
C Y
B Z
B Z
B X
B Z
C Y
B Y
B Z
B Z
B Y
B Z
A X
B X
B Y
A X
C Y
B Z
A Y
B Z
B Z
B Z
A X
B Y
B Z
B Y
B Z
A X
B Z
A X
B Y
B Y
B X
B X
A Y
B Z
C Y
A X
A X
B X
B Z
B X
B X
B Z
B Z
A X
B X
B Z
A Z
B Y
A X
B Z
C Y
B Z
B X
B Z
B Z
A Y
A Y
B Z
B X
B Y
B Z
A Y
A Y
B X
B X
C X
B Z
C Z
B X
A Y
B Z
A Z
B Y
A Y
B Z
B X
B X
A Y
B Z
B Y
A X
B Y
B Y
B Y
B X
B Y
B Z
B Y
B Y
A X
C Z
B Z
B Z
A Z
B Z
B Z
A X
B X
A Y
A X
B X
C X
B X
B Z
B Y
A Z
A Y
B Y
B Z
B X
B X
B Z
B Z
B Y
A X
B Z
A Y
A X
B Y
B Z
B Z
B X
A X
A Y
C X
A Y
B Z
B Y
B Z
B Z
C Z
B Z
B X
A X
C Y
B X
B Z
A Z
A X
A Y
B Z
B Z
A Y
A Y
B X
A Y
A Y
B Z
A Y
B Y
B Z
A Y
A Z
B X
B X
B Y
B Z
B Z
A X
B X
C X
B Y
A Y
A X
A X
A X
B Y
A X
A Z
A Y
B X
A Y
B Z
C X
B X
B X
B Z
B X
B X
B Z
A X
B Z
B X
B Z
A Y
B Y
B Z
B Y
B Z
B Z
A X
B Z
A Y
C Z
A Y
B Z
A Y
B X
B Z
B X
C Y
A X
B Z
B Y
A X
A X
B X
B Z
A Z
B Z
B X
B X
B X
B Z
B Z
B Y
B X
B X
A X
B Z
B X
A X
A X
A X
B X
B Z
A Z
B X
B Y
B Y
B Z
B Z
C Z
A Y
A X
B Y
B X
B Z
B X
A X
B Z
B Z
B Y
B Z
B X
B Z
B Z
B X
B X
A X
A X
B X
B Z
B Z
C X
C X
B Z
B Z
B X
B Z
B Y
B X
A X
B Y
A Y
A X
B Z
B Z
C X
B Y
B Z
B Z
B X
B Z
B Z
C Z
B Z
B Y
B Z
B X
A X
B X
B Z
B Z
A Y
B Z
B Z
B Z
B Z
A X
B Z
B Z
B X
B Y
B Y
B Z
A X
B Z
B Y
B X
A X
B X
B Z
B Z
B Z
B X
B Z
B Z
B Z
A Y
B Z
C X
B Y
B Y
B Y
A Y
B Z
B Z
A X
C Z
B Z
B Z
B Z
B X
B Z
A Z
B Z
B Z
B Z
A Z
B Z
C Y
B X
A X
A Y
B Z
B Z
A Z
B Z
B X
A Z
B Y
B Z
B X
B Z
B X
B Y
B Z
A X
B Z
B Z
A Y
B X
B X
B X
A Z
C X
A Z
B Z
B Z
B Z
B X
A Y
C X
A Z
A Y
B X
B Z
B X
B Z
B Y
A Y
B X
C X
A Y
C Z
A X
B Z
B Z
A Z
B X
B Y
B Z
A Y
B Y
A X
A X
C Y
A Y
B X
A X
B Y
B X
B Y
A Y
A X
C Y
B Y
B Y
B Z
B Y
B X
B Z
B Z
B X
A X
B Z
B Z
B Z
B X
B Z
B X
B Z
B Z
B Y
B Y
B X
A X
B Z
B Y
A Y
B Z
B X
B Z
B Z
A Z
B Z
A Y
B Z
A X
B Z
B Z
A Y
B Z
A X
B Z
A Y
A Y
A Z
B X
B Z
B Y
A Z
C Z
B Z
A X
A X
B X
A Z
B X
B X
B Z
C X
B Z
B Z
B Z
B Z
B X
A Z
A Y
B Z
B Y
C Z
B Y
B Z
A Y
B X
B X
B Z
A X
A Y
B Z
B X
B Y
A Y
C Y
C Y
B Z
A Y
B Y
A Y
B Z
B Z
A X
B X
A X
B X
A Y
A X
B Z
A X
B Z
B X
B Z
B X
B X
A X
A Y
B Z
B X
B X
A X
A Y
A X
B Y
B Z
B Z
B Z
B Z
B Z
B X
B Y
A Y
B X
B Z
A Y
B X
A X
B Z
C X
B Y
A Y
A X
A X
B X
B X
B Z
B Z
B Z
B X
B Y
B X
B Z
B X
B Y
B X
B Z
B X
B Y
B X
B Z
B Z
B Z
A X
C X
C Y
A Y
B X
B X
A Y
B Z
B X
B Z
B X
B X
C X
B Z
B Z
B Y
A Z
A Z
C X
B X
C Y
B Z
C X
B Z
A Y
C Z
B X
B Y
A Y
B Y
B X
B Z
A Y
A Z
C X
B Z
A X
B X
B Z
C Z
A Y
B Z
B Z
A X
A X
B Z
B Y
C Z
B Z
B X
B Z
A X
A X
B Y
A X
B X
A Y
B Z
A Y
B Z
A Y
B Z
A X
B Y
B Z
B Z
B Z
A X
A X
B X
B Z
A X
B Z
A X
A X
B Z
B Z
B Z
B X
B Y
B Z
B Z
B X
B Z
B Y
B Z
C Y
B X
C Z
B Z
B Z
A Y
B Z
B X
A Y
B Z
B X
B Z
B Z
B Y
B Z
A Z
A Y
B Z
B Z
B Y
A Y
B Z
A X
B X
A Z
A X
B X
B Z
B X
B Z
B Z
B Y
B Y
B Z
B Y
B Z
B Z
A X
B X
B Y
B Y
C Z
A X
B Y
A Y
B Z
B Z
B Z
C X
B X
A Z
B Y
A X
C X
B Z
B Z
B Z
B Y
A Y
A Y
B Y
B Z
B Z
B Z
C X
A X
B Z
A Y
B X
B Z
B Z
B Z
B Z
B X
B Y
B Z
B Z
A Y
C X
A X
B Z
B Z
A Y
A Y
B Z
A Y
B Y
A X
B Z
B Z
A Y
B Z
B Z
A X
A X
B X
B Z
B Z
A X
B Z
B Z
C Y
B Z
A X
B Z
A X
B Z
A X
B X
A X
A Y
B Z
B X
A X
B Z
A X
A Z
B Z
B Z
B Z
A Y
B X
A X
B Y
A Y
B Z
B Z
B X
B Z
B X
B X
A X
B Z
A Y
A X
B X
A Z
B X
B Y
B Z
B X
B X
B Z
C X
A X
B Z
B Y
C Z
B Z
A Z
A Z
A X
A Y
B Z
B Z
B X
A Z
B Z
B Z
B X
B Y
B X
B Z
B X
B Z
A Y
A X
B Z
B X
B Y
B Z
B Z
B Z
B Z
C X
C X
B Z
B X
B Z
B Z
B Z
B X
B Z
B Z
B Z
B X
B Z
B Z
B Z
C Y
B Z
B Y
B Z
A Z
A Y
B Z
A Y
C X
A X
B X
A Y
B Z
A Y
B Z
B Y
B Z
B Z
C Y
B Y
B Z
B Z
B X
B Y
B Z
B Y
B Z
B X
B Y
B Z
B Z
B Z
B Y
B Y
B Z
B Z
C X
B Z
A Y
B Z
B Z
B Y
B Y
B Z
B Z
A Y
B Z
A Z
C X
A Y
A Y
A X
B Z
A X
C Y
A Z
C Y
C X
B Z
A X
A Y
B Z
B Z
B Z
B Z
B Y
A Z
A Y
B Y
A Y
A Z
B X
B Z
B Z
A X
C Y
B Z
B X
C X
A Z
B Z
B X
B Y
A X
A X
B Z
C X
B Z
B Z
B X
B Z
B Z
B Z
A X
B Z
B Y
B Z
B Z
B Y
B X
A X
B Z
A X
A X
B Y
B Z
B X
A Y
C X
B Y
A X
A X
A X
C Y
B X
C Z
A X
B Z
A Z
B Z
A X
A Y
B Z
A X
B X
A Y
A X
A Y
B X
B Y
B Z
B Y
C X
C Y
B Z
B Y
B Z
A X
C Z
B Z
A X
B X
A X
B X
B Z
B X
B Z
B Z
B Z
A X
A X
B Z
B X
B Z
B Z
B Z
A X
B Z
B X
A X
C Z
A Y
B Z
A Y
B Z
B Z
B Z
C Z
B Z
B Z
A X
B X
A X
B X
A X
B Z
B Z
B X
B Z
A X
A Y
A Y
A Y
B Z
B Z
B Z
B Z
A X
B Z
B Z
B Y
B Z
A Z
B Z
A Y
B Y
B Y
C Y
B X
B Z
B Y
B Z
B Y
B Z
B Z
A Y
B X
A Z
B Z
A X
B Y
A X
B Z
B Y
A X
B X
B Z
B Z
B Z
A Y
A Y
A X
A Y
A Y
B X
B Z
B Y
B Y
B X
B Y
B X
B X
B X
B Z
B Z
B Z
A Y
A Z
B Y
C X
B X
B Z
C Y
B Z
C Y
B Z
B Z
B Z
B X
A X
B Z
B Z
A X
B Z
B Z
B Z
B Z
B Z
B Z
B X
A X
A Y
B Y
B Z
B Y
B X
A Z
A X
B X
B Y
B Z
B Z
B Z
B X
C X
B X
A Z
A X
A Y
B Z
B Z
A X
B Z
A Y
B X
B Z
B Y
B Y
A X
B X
B Z
B Z
B Z
B Z
B Y
A Z
A Y
B X
A X
B Y
B Z
C Y
B Z
B Z
B X
A Z
A Y
B Y
B Z
A Y
B Y
A Y
A X
A Y
B Y
C X
C Z
B X
A Z
A X
B Y
B X
A X
B Z
A Y
A Z
B Z
B Z
B Z
A X
B Z
B Y
B Z
A X
A Y
B Z
B X
C X
B Z
A Y
B Z
B Z
B X
B Z
B Z
A Y
A X
B Z
C Y
A Y
B Y
B Z
A X
B Z
B Z
B Z
B Z
C X
B Y
B Z
B X
B Z
B Z
B Z
B Z
B Z
B Y
A X
B Y
A Y
A X
A X
B Z
B Z
B Z
C X
B X
B Z
A X
A X
B Y
A X
B Z
B Z
B X
B Y
B Y
B Y
A Y
A Y
A X
A X
B Z
B Y
B Y
B X
B Y
B Z
A Z
B Z
A Z
B X
B Y
C Y
B X
B X
A X
A X
A X
A X
B Z
B Z
A Z
B X
B X
B X
B X
B X
B Z
B Z
B Z
B Z
B Z
B Y
B Z
B Z
A Y
A Z
C Y
B Y
A X
B Z
B X
A Y
B Y
A Y
C Z
A X
A Y
B X
B X
C Z
B Y
A Y
A Y
A X
B X
A X
B X
A Y
B Z
B Z
A Z
B X
B Z
B Z
B Y
A Y
B Z
B Y
B Z
B Z
B Z
C X
B Y
A Y
B Z
A Y
A Y
A X
B X
B Y
C Y
C X
B Z
B Z
B Z
A Y
A Z
B Y
B X
B X
B Y
B Z
B Z
A Y
B Y
B Z
B Z
B Z
A X
C Y
B Z
A Z
C Z
B X
B Y
B Z
A Z
A X
B X
A X
C X
B Z
A X
B Y
C Y
B Y
B Z
B X
A Y
B Z
B Z
B Z
A X
B Z
B Z
B Z
B Z
A X
B Z
B X
B X
B Z
C X
A Y
B Z
B X
B Z
B Y
B Z
B Y
C Y
A X
B Y
B Z
B Z
B X
B Z
A X
B Z
A X
B Y
B Z
A X
B Z
B X
A Z
B Z
C X
B Z
B Y
B Y
B Y
B Z
B X
B X
A X
A X
B Z
B Z
A X
A X
A Z
A Y
C Z
B Z
B Y
C X
B X
B Y
B Y
A Y
B Z
B Z
B Y
C Y
B Z
A Y
B Z
B Z
B Z
B Z
A Y
B Z
B Z
B Z
A X
B X
A Y
B Y
B Z
B Z
B Z
A Y
B Z
B Z
B Z
B Y
B Y
B X
B X
B Y
C Z
B X
B Y
C Y
A Y
B Z
A Z
B Z
A Y
B X
B Z
B X
C Z
A Y
B X
B Z
A X
A X
B Z
B Z
B Y
B Z
C Y
C X
B Z
B Z
A X
B X
A Y
B Y
B Y
C Z
A X
B X
B Z
A Y
B Z
C X
B Z
A Z
C X
A Y
A Z
B Z
A Y
A X
B Z
B Y
B Z
B X
B Z
A X
B Z
A Z
B Z
B Z
B Y
A X
B Z
B X
B Y
B Y
A Y
A X
C Z
B Y
C Z
A X
B X
B Z
B X
C X
B Z
B Z
A X
A Y
A Z
A Z
B Z
B Z
A X
B Z
B Y
A X
A Y
B Z
B X
A Y
B Z
B Y
B X
B Z
A Z
B Y
A Y
A X
B Z
B Z
B Z
B Z
A X
B Y
B Z
B Z
B Z
A Z
B Y
A X
C Y
B Z
B Z
B Y
B X
B Z
A Y
A X
B Z
B Y
B Z
A Y
C X
B Y
B Z
B Z
A Y
B X
B Z
B Z
A X
B X
B Z
A Z
B Y
B Z
B Z
B Z
B Z
B Z
A Z
B X
A Y
C Y
B Z
A Y
B Z
B Z
A Y
B Z
B Z
A Y
A Y
B X
B Y
A X
B Z
A Y
B Z
A Z
B Z
B X
B Z
B Z
B Y
B Y
B Z
A X
B Y
B Z
B Z
B Z
B Y
A Y
B Z
B Y
B Z
B Z
B X
A X
B Y
A X
A X
B Z
A X
B Z
C X
B Y
B Z
B Z
B Y
A X
B Z
B Z
B Y
B X
B Z
B X
B Z
C Y
B Y
B Z
B Z
A Z
B Z
B Z
A Y
B Z
B Y
A X
B Z
B Z
B Z
B X
B Z
B Y
B Z
B X
C Z
B Y
B Z
B Z
B Y
A X
A X
B Y
A Y
B Z
B X
B Y
B Z
B Y
B Z
A X
C X
B Z
A Z
C Y
A X
B Y
B Y
A X
A X
B Y
B Z
C Y
B Z
B Y
A X
B Y
B Y
B Z
B X
B Z
C X
C Y
B X
A Y
B Z
B Z
B Z
A X
B Z
B Z
B X
B X
B X
B X
B Y
A Y
B Z
B Z
B Z
A Z
B X
B X
B X
B X
B Z
A Y
B Y
B Z
A X
A X
A X
B Y
B Z
B X
B Z
A Z
B Z
B Z
A Y
C Z
A X
B Z
C X
B Z
B Z
B Z
B Y
B Y
B Z
C X
B Z
A Y
B X
C X
A Y
C X
B Z
B X
A Z
B Z
C Y
B Y
B X
B Z
B Y
A Y
B X
A Z
B Z
B Z
C Z
B Z
B Y
B Z
A Y
C Y
B Z
B X
A Y
B Z
B Z
B Y
B Y
B Z
A X
A Y
A X
A Z
B Z
B Z
B Y
A Y
B Y
B Z
B Y
A Y
B X
A X
B Y
B Y
B X
B Z
B Y
C X
B Z
B Z
B Z
A X
B Z
B X
A Y
A Y
B Z
A Y
A Y
A Y
B Y
B Z
A X
C X
B Y
A X
C X
B Z
B Y
A Y
B X
B Z
B Y
B Y
B Z
B Z
A Y
B X
B X
A X
B Z
B X
B Z
B Z
C Z
C Y
A X
B X
C X
B X
A Z
A X
B Y
C X
B Y
A X
A X
A X
B X
B Z
B X
B Y
B Z
B Z
B Z
B Z
A X
A X
B X
B Z
A Z
A X
B Z
C Y
B Z
A X
B Z
A Z
B Z
B Z
A Y
B Y
B Z
B X
B Z
B Z
C Y
B Z
B X
A X
B Y
B Z
A Y
A X
B Z
B Y
B Z
B X
A Y
C Z
B X
B Z
A X
B X
B Z
B Z
A X
B Y
B X
B Z
B Z
B Y
B X
B X
C X
B Z
C X
B Z
C Y
A Y
C X
B X
B Y
B Z
B Y
B Z
A X
B Z
B Z
B Y
B Y
A X
A X
A Y
B Z
B Z
A Y
B Z
A X
B Z
B X
B X
B X
A Y
B Z
B Z
A Y
B Z
A Z
B Y
B X
A Z
A Z
B Z
A Y
B Z
B Y
B Y
B X
A Y
B X
B Z
B Z
B Y
B X
B Z
B Z
B Z
B Z
A X
B X
B X
B Z
A X
B X
B Z
B Z
B Y
A Y
B Z
B Y
B Y
A X
B X
B Y
C X
C X
B Y
A Y
A Z
B X
B Z
B Z
B Z
B Z
A X
A Y
A Y
B Y
C X
B Z
A Y
B X
C Z
B X
B Y
A Y
B X
B Y
B Z
B Y
B X
B Y
A Z
B Z
A X
B X
B X
B Z
B Y
B Y
B Y
A X
B X
B Y
B Z
A Y
A X
B X
B Z
B Z
C X
A Z
B Z
A X
A Y
C X
B Z
B Z
A X
B X
A Y
A Y
B Z
B Z
B X
A Z
B Z
B Y
A Y
B Z
B Y
A Y
B Y
B Z
A X
B Y
A X
A Y
A X
A X
B Y
A Z
B Z
A Y
B Z
A Z
A Y
B Z
B Y
B Z
B X`
