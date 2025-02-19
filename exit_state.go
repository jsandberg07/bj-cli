package main

func (gs *Gamestate) GetExitState() *State {
	s := State{
		Logic: ExitLogic,
		Print: ExitPrint,
	}

	return &s
}

func ExitLogic(gs *Gamestate) {
	gs.Playing = false
}

func ExitPrint(gs *Gamestate) {
	// nothing again
}
