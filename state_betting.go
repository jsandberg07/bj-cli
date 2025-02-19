package main

import (
	"fmt"
	"strconv"
)

func (gs *Gamestate) getBettingState() *State {
	s := State{
		Logic: BettingLogic,
		Print: BettingPrint,
	}
	return &s
}

func BettingLogic(gs *Gamestate) {
	fmt.Printf("You have $%v\n", gs.P.Money)
	fmt.Println("Enter amount to bet or 'exit' to quit")
	for {
		input, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if input == "exit" {
			gs.SetNextState(gs.GetExitState())
			return
		}
		bet, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if bet <= 0 {
			fmt.Println("Can't wager nothing")
			continue
		}
		if bet > gs.P.Money {
			fmt.Println("Can't wager more than you have")
			continue
		}

		gs.P.Bet = bet
		break
	}

	gs.SetNextState(gs.GetPlayingState())
}

func BettingPrint(gs *Gamestate) {
	// TODO: add checks to make sure things are actually set lmao
	fmt.Println("Set!")
}
