package main

import (
	"cmp"
	"fmt"
)

const scoreTarget = 21

type Result int

const (
	ResultWin Result = iota
	ResultDraw
	ResultLose
	ResultError
)

// blackjack pays 3 to 2

type Gamestate struct {
	Player Player
	Dealer Dealer
	Deck   Deck
}

func (gs *Gamestate) Init() {
	gs.Player = Player{}
	gs.Player.Init("Player")

	gs.Dealer = Dealer{}
	gs.Dealer.Init("Dealer")

	gs.Deck = Deck{}
	gs.Deck.Init()
}

func (gs *Gamestate) Reset() {
	gs.Player.Reset()
	gs.Dealer.Reset()
	gs.Deck.Shuffle()
}

// how do we format it
// figure that out, and how to transplant strings
func (gs *Gamestate) Print() string {
	ps := gs.Player.Print()
	ds := gs.Dealer.Print()
	return printPlayers([][]string{ps, ds})
}

func (gs *Gamestate) CompareHands() Result {
	switch cmp.Compare[int](gs.Player.Hand.Score, gs.Dealer.Hand.Score) {
	case -1:
		return ResultLose
	case 0:
		return ResultDraw
	case 1:
		return ResultWin
	default:
		fmt.Println("Default in CompareHands")
		return ResultError
	}
}

func (gs *Gamestate) PlayRound() Result {
	// deal two cards to the player
	// deal two cards to the dealer
	for i := 0; i < 2; i++ {
		gs.Player.TakeCard(gs.Deck.Deal())
		gs.Dealer.TakeCard(gs.Deck.Deal())
	}
	// print
	fmt.Print(gs.Print())
	// ask the player if he wants to hit or stand
	playerStand := false
	for {
		fmt.Printf("Score: %v\n", gs.Player.GetScore())
		cmd := gs.Player.GetPlayerChoice()
		switch cmd {
		case CommandHit:
			gs.Player.TakeCard(gs.Deck.Deal())
		case CommandStand:
			playerStand = true
		default:
			fmt.Println("Default found in Play")
			playerStand = true
		}

		fmt.Print(gs.Print())

		if gs.Player.IsBust() {
			return ResultLose
		}
		if playerStand {
			break
		}
	}
	// if hit, deal card, calc bust, check if lose
	dealerStand := false
	for {
		cmd := gs.Dealer.MakeChoice()
		switch cmd {
		case CommandHit:
			gs.Dealer.TakeCard(gs.Deck.Deal())
		case CommandStand:
			dealerStand = true
		default:
			fmt.Println("Default found in Play")
			dealerStand = true
		}
		if gs.Dealer.IsBust() {
			return ResultWin
		}
		if dealerStand {
			break
		}

	}
	fmt.Print(gs.Print())
	fmt.Printf("Score: %v\n", gs.Player.GetScore())

	return gs.CompareHands()
}
