package main

import (
	"fmt"
	"strconv"
)

func (gs *Gamestate) getBettingState() *State {
	s := State{
		Logic: BettingLogic,
		// Print: BettingPrint,
	}
	return &s
}

func BettingLogic(gs *Gamestate) {
	if gs.Player.Money == 0 {
		fmt.Println("You're out of money...")
		gs.SetNextState(gs.GetExitState())
	}
	fmt.Printf("You have $%v\n", gs.Player.Money)
	fmt.Println("Enter amount to bet, 'save', or 'exit'")
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

		case "save":
			err := gs.Save()
			if err != nil {
				fmt.Println("Error saving")
				fmt.Println(err)
			}
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
