package main

import (
	"fmt"
	"strconv"
)

func (gs *Gamestate) getMainMenuState() *State {
	s := State{
		Logic: mainMenuLogic,
		Print: mainMenuPrint,
	}
	return &s
}

func mainMenuLogic(gs *Gamestate) {
	fmt.Println("Enter player name")
	for {
		name, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if name == "" {
			fmt.Println("No name found")
			continue
		} else {
			gs.Player.Init(name)
			break
		}
	}

	fmt.Println("Enter number of bots")
	for {
		num, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if num == "" {
			gs.AddBots(0)
			break
		}
		bots, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gs.AddBots(bots)
		break
	}

	fmt.Println("Enter number of decks used by dealer")
	for {
		num, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if num == "" {
			gs.NumDecks = 1
			break
		}
		decks, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gs.NumDecks = decks
		break
	}

	fmt.Println("Enter goal for the player, or nothing to have no goal")
	for {
		num, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if num == "" {
			gs.Player.Goal = 0
			gs.Player.GoalMet = true
			break
		}
		goal, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gs.Player.Goal = goal
		gs.Player.GoalMet = false
		break
	}

	gs.SetNextState(gs.getNewGameState())
}

func mainMenuPrint(gs *Gamestate) {
	// literally does nothing
}
