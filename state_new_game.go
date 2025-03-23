package main

func (gs *Gamestate) getNewGameState() *State {
	s := State{
		Logic: mainNewGameLogic,
		// Print: mainNewGamePrint,
	}
	return &s
}

func mainNewGameLogic(gs *Gamestate) {
	gs.Dealer.Init("Dealer")
	gs.Deck.Init(gs.NumDecks)
	gs.SetNextState(gs.getBettingState())
}
