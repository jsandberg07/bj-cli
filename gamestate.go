package main

import "fmt"

const scoreTarget = 21

// blackjack pays out 3/2 (or can be changed to 6/5 or something)
const blackjackPayout float64 = (3.0 / 2.0)

type Result int

const (
	ResultWin Result = iota
	ResultDraw
	ResultLose
	ResultBlackjack
	ResultSurrender
	ResultError
)

type State struct {
	Logic func(gs *Gamestate)
	// removed because it was anticipated to be used but most functions were just empty
	// Print func(gs *Gamestate)
}

type Gamestate struct {
	S        *State
	NS       *State
	Player   Player
	Dealer   Dealer
	Bots     []Bot
	Deck     Deck
	NumDecks int
	Playing  bool
}

func (gs *Gamestate) Logic() {
	gs.S.Logic(gs)
}

func (gs *Gamestate) Init() {
	// set variables
	gs.Playing = true
	gs.Player = Player{}
	gs.Dealer = Dealer{}
	gs.S = gs.getMainMenuState()
}

func (gs *Gamestate) Reset() {
	gs.Player.Reset(1)
	gs.Dealer.Reset()
	gs.Deck.Shuffle()
}

func (gs *Gamestate) Run() {
	for {
		gs.Logic()
		gs.CheckNextState()
		if !gs.Playing {
			return
		}
	}

}

func (gs *Gamestate) PrintTable() string {
	p := [][]string{}
	p = append(p, gs.Player.Print())
	p = append(p, gs.Dealer.Print())
	for _, b := range gs.Bots {
		p = append(p, b.Print())
	}

	return printPlayers(p)
}

func (gs *Gamestate) SetNextState(s *State) {
	// fmt.Println()
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
		gs.Player.Probability.RemoveCard(card.GetValue())
	}
	return card
}

// string return is just to break up visuals
func (gs *Gamestate) FlipCards() string {
	for i := 0; i < len(gs.Dealer.Hand.Cards); i++ {
		gs.Dealer.Hand.Cards[i].Visible = VisibleFaceup
	}
	for i := 0; i < len(gs.Bots); i++ {
		for j := 0; j < len(gs.Bots[i].Hand.Cards); j++ {
			gs.Bots[i].Hand.Cards[j].Visible = VisibleFaceup
		}
	}
	return "Flip!"
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
	gs.Player.Reset(1)
	gs.Dealer.Reset()
	for i := 0; i < len(gs.Bots); i++ {
		gs.Bots[i].Reset()
	}
}

func (gs *Gamestate) AddBots(num int) {
	if num == 0 {
		return
	}

	for i := 0; i < num; i++ {
		bot := Bot{}
		bot.Init(fmt.Sprintf("Bot %v", i+1))
		gs.Bots = append(gs.Bots, bot)
	}
}
