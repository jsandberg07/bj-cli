package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (gs *Gamestate) getMainMenuState() *State {
	s := State{
		Logic: mainMenuLogic,
		Print: mainMenuPrint,
	}
	return &s
}

func mainMenuLogic(gs *Gamestate) {
	fmt.Println("Enter player name. Enter name of save file to load")
	for {
		name, err := getInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if name == "" {
			fmt.Println("No name found")
			continue
		}

		// try to load. if file not found, start new game with name
		err = gs.Load(name)
		if err != nil && !strings.Contains(err.Error(), "no such file or directory") {
			// any other error other than file not being found, exit
			fmt.Println(err)
			gs.SetNextState(gs.GetExitState())
			return
		}
		// file found and load successful
		if err == nil {
			err := gs.CleanSave()
			if err != nil {
				fmt.Println("Error cleaning save file")
				fmt.Println(err)
			}
			gs.SetNextState(gs.getNewGameState())
			return
		}

		gs.Player.Init(name)
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
			gs.Player.Stats.SetGoal(0)
			break
		}
		goal, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		gs.Player.Stats.SetGoal(goal)
		break
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

	gs.SetNextState(gs.getNewGameState())
}

func mainMenuPrint(gs *Gamestate) {
	// literally does nothing
}
