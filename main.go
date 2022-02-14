package main

import (
	"fmt"
	"github.com/gramm500/rockPaperScissors/entity"
	"os"
)

func main() {
	enemyFigureStr := os.Args[1]
	enemyFigure, err := entity.GenerateFigureFromArg(enemyFigureStr)
	if err != nil {
		return
	}
	sc, err := entity.GenerateRandomFigure()
	if err != nil {
		fmt.Println("ERROR GENERATING FIGURE")
		return
	}
	fmt.Println("YOUR FIGURE", enemyFigure.GetValue())
	fmt.Println("GENERATED FIGURE", sc.GetValue())
	entity.FightResult(sc, enemyFigure)
}
