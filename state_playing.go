package main

import (
	"cmp"
	"fmt"
	"os"
)

func (gs *Gamestate) GetPlayingState() *State {
	s := State{
		Logic: PlayingLogic,
		Print: PlayingPrint,
	}
	return &s
}

func PlayingLogic(gs *Gamestate) {
	result := gs.PlayRound()
	switch result {
	case ResultWin:
		gs.P.Win()
	case ResultDraw:
		gs.P.Draw()
	case ResultLose:
		gs.P.Lose()
	default:
		fmt.Println("Default in playing logic!")
		os.Exit(1)
	}

	gs.Cleanup()

	gs.SetNextState(gs.getBettingState())

}

func PlayingPrint(gs *Gamestate) {

}

func (gs *Gamestate) PlayRound() Result {
	// deal two cards to the player
	// deal two cards to the dealer
	for i := 0; i < 2; i++ {
		gs.P.TakeCard(gs.Deck.Deal())
		gs.D.TakeCard(gs.Deck.Deal())
	}
	// print
	fmt.Print(gs.Print())
	// ask the player if he wants to hit or stand
	playerStand := false
	for {
		fmt.Printf("Score: %v\n", gs.P.Hand.Score)
		cmd := gs.P.GetPlayerChoice()
		switch cmd {
		case CommandHit:
			gs.P.TakeCard(gs.Deck.Deal())
		case CommandStand:
			playerStand = true
		default:
			fmt.Println("Default found in Play")
			playerStand = true
		}

		fmt.Print(gs.Print())

		if gs.P.IsBust() {
			fmt.Printf("Bust! - %v\n", gs.P.Hand.Score)
			return ResultLose
		}
		if playerStand {
			break
		}
	}
	// if hit, deal card, calc bust, check if lose
	dealerStand := false
	for {
		cmd := gs.D.MakeChoice()
		switch cmd {
		case CommandHit:
			gs.D.TakeCard(gs.Deck.Deal())
		case CommandStand:
			dealerStand = true
		default:
			fmt.Println("Default found in Play")
			dealerStand = true
		}
		if gs.D.IsBust() {
			return ResultWin
		}
		if dealerStand {
			break
		}

	}
	fmt.Print(gs.Print())
	fmt.Printf("Score: %v\n", gs.P.Hand.Score)

	return gs.CompareHands()
}

func (gs *Gamestate) CompareHands() Result {
	fmt.Printf("%v - %v\n", gs.P.Hand.Score, gs.D.Hand.Score)
	switch cmp.Compare[int](gs.P.Hand.Score, gs.D.Hand.Score) {
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
