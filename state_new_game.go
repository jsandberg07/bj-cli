package main

import (
	"fmt"
)

func (gs *Gamestate) getNewGameState() *State {
	s := State{
		Logic: mainNewGameLogic,
		Print: mainNewGamePrint,
	}
	return &s
}

func mainNewGameLogic(gs *Gamestate) {
	gs.Dealer.Init("Dealer")
	gs.Deck.Init()
	gs.SetNextState(gs.getBettingState())
}

func mainNewGamePrint(gs *Gamestate) {
	// TODO: add checks to make sure things are actually set lmao
	fmt.Println("Set!")
}
