package main

import "fmt"

func (gs *Gamestate) getMainMenuState() *State {
	s := State{
		Logic: mainMenuLogic,
		Print: mainMenuPrint,
	}
	return &s
}

func mainMenuLogic(gs *Gamestate) {
	fmt.Println("Enter player name")
	var name string
	var err error
	for {
		name, err = getInput()
		if err != nil {
			fmt.Println(err)
		}
		if name == "" {
			fmt.Println("No name found")
			continue
		} else {
			gs.P.Init(name)
			break
		}
	}

	gs.SetNextState(gs.getNewGameState())
}

func mainMenuPrint(gs *Gamestate) {
	// literally does nothing
}
