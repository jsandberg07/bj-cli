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

	gs.SetNextState(gs.getNewGameState())
}

func mainMenuPrint(gs *Gamestate) {
	// literally does nothing
}
