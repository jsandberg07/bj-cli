package main

import (
	"fmt"
)

func main() {
	fmt.Println("Here we go")
	gs := Gamestate{}
	gs.Init()
	result := gs.PlayRound()
	switch result {
	case ResultWin:
		fmt.Println("You win!")
	case ResultDraw:
		fmt.Println("You draw!")
	case ResultLose:
		fmt.Println("You lose!")
	case ResultError:
		fmt.Println("Result error!")
	default:
		fmt.Println("Default in results")
	}
}
