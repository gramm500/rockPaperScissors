package entity

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const DRAW = 2
const WIN = 1
const LOSE = 0

var FightResultMap = map[int]string{
	0: "Lose", 1: "Win", 2: "Draw",
}

var FightMap = map[string]map[string]int{
	"rock":     {"rock": DRAW, "scissors": WIN, "paper": LOSE},
	"paper":    {"rock": WIN, "scissors": LOSE, "paper": DRAW},
	"scissors": {"rock": LOSE, "scissors": DRAW, "paper": WIN},
}

type Turn struct {
	PreviousFigure *Figure
	PreviousTurn   bool
}

func NewTurn(prevFigure *Figure, prevTurn bool) Turn {
	return Turn{
		PreviousFigure: prevFigure,
		PreviousTurn:   prevTurn,
	}
}

func GenerateRandomFigure() (Figure, error) {
	rand.Seed(time.Now().UnixNano())
	val := rand.Intn(3)

	if val == 0 {
		return NewPaper(), nil
	}

	if val == 1 {
		return NewScissors(), nil
	}

	if val == 2 {
		return NewRock(), nil
	}

	return nil, errors.New("NO FIGURE")
}

func GenerateFigureFromArg(arg string) (Figure, error) {
	if strings.ToLower(arg) == "paper" {
		return NewPaper(), nil
	}

	if strings.ToLower(arg) == "scissors" {
		return NewScissors(), nil
	}

	if strings.ToLower(arg) == "rock" {
		return NewRock(), nil
	}

	return nil, errors.New("no such figure found")
}

func FightResult(botFigure Figure, enemyFigure Figure) {
	fightres, _ := FightMap[enemyFigure.GetValue()][botFigure.GetValue()]
	fmt.Println(FightResultMap[fightres])
}
