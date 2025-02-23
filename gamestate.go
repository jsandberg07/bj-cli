package main

import "fmt"

const scoreTarget = 21

type Result int

const (
	ResultWin Result = iota
	ResultDraw
	ResultLose
	ResultError
)

// blackjack pays 3 to 2

// state needs properties
// logic
// just use state enum and a switch
// everything else can be kept and changed as needed
// and the state just points to the logic and render

type State struct {
	Logic func(gs *Gamestate)
	Print func(gs *Gamestate)
}

type Gamestate struct {
	S       *State
	NS      *State
	P       Player
	D       Dealer
	Deck    Deck
	Playing bool
}

func (gs *Gamestate) Logic() {
	gs.S.Logic(gs)
}

func (gs *Gamestate) Init() {
	// set variables
	gs.Playing = true
	gs.P = Player{}
	gs.D = Dealer{}
	gs.S = gs.getMainMenuState()
}

func (gs *Gamestate) Reset() {
	gs.P.Reset(1)
	gs.D.Reset()
	gs.Deck.Shuffle()
}

func (gs *Gamestate) Run() {
	for {
		gs.Logic()
		gs.Print()
		gs.CheckNextState()
		if !gs.Playing {
			return
		}
	}

}

// how do we format it
// figure that out, and how to transplant strings
func (gs *Gamestate) Print() string {
	ps := gs.P.Print()
	ds := gs.D.Print()
	return printPlayers([][]string{ps, ds})
}

func (gs *Gamestate) SetNextState(s *State) {
	fmt.Println()
	gs.NS = s
}

func (gs *Gamestate) CheckNextState() {
	if gs.NS != nil {
		gs.S = gs.NS
		gs.NS = nil
	}
}

func (gs *Gamestate) Deal(v Visible) Card {
	card := gs.Deck.Deal(v)
	if v == VisibleFaceup {
		gs.P.Probability.RemoveCard(card.GetValue())
	}
	return card
}

// OOP was a mistake
func (gs *Gamestate) FlipCards() {
	for i := 0; i < len(gs.D.Hand.Cards); i++ {
		gs.D.Hand.Cards[i].Visible = VisibleFaceup
	}
}

// lets create some states
// without knowing what to do lmao
// so we have a state
// we have logic
// we render
// the state is state is a command.logic
// and .render which we pretty much have
// i want a start menu to save + load + new game
// playing the game
// and a menu for saving + betting + quitting
// problem is state making changes to base state functions
// so we have to transfer them
// imagine actually making a game

// so we have
// GS{Menu} logic -> get input and switch it, render -> is what you expect
// too bad go isn't OOP and doesnt have inhertiance
// but the basis is state->logic which pokes members then ->render which renders it
// how do we switch states?
// i guess create a new state and make it the next one, then swap
// but where do we store the important like deck before it's created

func (gs *Gamestate) Cleanup() {
	gs.Deck.Reset()
	gs.P.Reset(1)
	gs.D.Reset()
}
