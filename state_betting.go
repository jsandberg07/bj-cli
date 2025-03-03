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
	fmt.Printf("You have $%v\n", gs.Player.Money)
	fmt.Println("Enter amount to bet or 'exit' to quit")
	for {
		input, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// try to handle as an int for a bet first
		bet, err := strconv.Atoi(input)
		if err == nil {
			if bet <= 0 {
				fmt.Println("Can't wager nothing")
				continue
			}
			if bet > gs.Player.Money {
				fmt.Println("Can't wager more than you have")
				continue
			}

			gs.Player.Bet = bet
			break
		}

		// if that doesn't work, try it as a word command
		switch input {
		case "exit":
			gs.SetNextState(gs.GetExitState())
			return

		case "stats":
			fmt.Println(gs.Player.Stats.PrintStats(gs.Player.Money))

		default:
			fmt.Println("Command not recognized")
		}
	}

	gs.SetNextState(gs.GetPlayingState())
}

func BettingPrint(gs *Gamestate) {
	// TODO: add checks to make sure things are actually set lmao
	fmt.Println("Set!")
}
