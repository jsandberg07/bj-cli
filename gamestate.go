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
	Player Hand
	Dealer Hand
	Deck   Deck
}

func (gs *Gamestate) Init() {
	gs.Player = Hand{}
	gs.Player.Init("Player")

	gs.Dealer = Hand{}
	gs.Dealer.Init("Dealer")

	gs.Deck = Deck{}
	gs.Deck.Init()
}

func (gs *Gamestate) Reset() {
	gs.Player.Reset()
	gs.Dealer.Reset()
	gs.Deck.Shuffle()
}

func (gs *Gamestate) Print() {
	gs.Dealer.PrintHand()
	gs.Player.PrintHand()
}

func (gs *Gamestate) CompareHands() Result {
	switch cmp.Compare[int](gs.Player.Score, gs.Dealer.Score) {
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
	gs.Print()
	// ask the player if he wants to hit or stand
	playerStand := false
	for {
		cmd := gs.Player.GetPlayerChoice()
		switch cmd {
		case CommandHit:
			gs.Player.TakeCard(gs.Deck.Deal())
			gs.Player.PrintHand()
		case CommandStand:
			playerStand = true
		default:
			fmt.Println("Default found in Play")
			playerStand = true
		}
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
		cmd := gs.Dealer.MakeDealerChoice()
		switch cmd {
		case CommandHit:
			gs.Dealer.TakeCard(gs.Deck.Deal())
			gs.Dealer.PrintHand()
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
	return gs.CompareHands()
}
