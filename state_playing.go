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
		gs.Player.Win()
	case ResultDraw:
		gs.Player.Draw()
	case ResultLose:
		gs.Player.Lose()
	default:
		fmt.Println("Default in playing logic!")
		os.Exit(1)
	}

	gs.Cleanup()

	gs.SetNextState(gs.getBettingState())

}

func PlayingPrint(gs *Gamestate) {

}

// TODO: encapsulate behavior in like player.turn() and bot.turn()
func (gs *Gamestate) PlayRound() Result {

	// deal two cards to the player
	for i := 0; i < 2; i++ {
		gs.Player.TakeCard(gs.Deal(VisibleFaceup))
	}
	// deal two cards to the dealer, one face up and one face down
	for i := 0; i < 1; i++ {
		gs.Dealer.TakeCard(gs.Deal(VisibleFaceup))
		gs.Dealer.TakeCard(gs.Deal(VisibleFacedown))
	}
	// deal two cards to the bots, one face up and one face down
	for i := 0; i < len(gs.Bots); i++ {
		gs.Bots[i].TakeCard(gs.Deal(VisibleFaceup))
		gs.Bots[i].TakeCard(gs.Deal(VisibleFacedown))
	}
	// print
	fmt.Print(gs.Print())
	// print probability
	toTarget, toBust := gs.Player.Probability.GetOdds(gs.Player.Hand.Score)
	fmt.Printf("To Target: %.3f - To Bust: %.3f\n", toTarget, toBust)
	// ask the player if he wants to hit or stand
	playerStand := false
	for {
		// player takes their turn
		fmt.Printf("Score: %v\n", gs.Player.Hand.Score)
		cmd := gs.Player.GetPlayerChoice()
		switch cmd {
		case CommandHit:
			gs.Player.TakeCard(gs.Deal(VisibleFaceup))
		case CommandStand:
			playerStand = true
		default:
			fmt.Println("Default found in Play")
			playerStand = true
		}

		// bots take their turn
		for i := 0; i < len(gs.Bots); i++ {
			if !gs.Bots[i].Standing {
				cmd := gs.Bots[i].MakeChoice()
				switch cmd {
				case CommandHit:
					gs.Bots[i].TakeCard(gs.Deal(VisibleFaceup))
				case CommandStand:
					gs.Bots[i].Standing = true
				default:
					fmt.Println("Default found in Play")
					gs.Bots[i].Standing = true
				}

				if gs.Bots[i].IsBust() {
					fmt.Printf("%s is bust!\n", gs.Bots[i].Name)
					gs.Bots[i].Standing = true
				}
			}
		}

		fmt.Print(gs.Print())

		if gs.Player.IsBust() {
			gs.FlipCards()
			fmt.Print(gs.Print())
			fmt.Printf("Bust! - %v\n", gs.Player.Hand.Score)
			return ResultLose
		}
		if playerStand {
			break
		}
		toTarget, toBust := gs.Player.Probability.GetOdds(gs.Player.Hand.Score)
		fmt.Printf("To Target: %.3f - To Bust: %.3f\n", toTarget, toBust)
	}
	// if hit, deal card, calc bust, check if lose
	dealerStand := false
	for {
		cmd := gs.Dealer.MakeChoice()
		switch cmd {
		case CommandHit:
			gs.Dealer.TakeCard(gs.Deal(VisibleFaceup))
		case CommandStand:
			dealerStand = true
		default:
			fmt.Println("Default found in Play")
			dealerStand = true
		}
		if gs.Dealer.IsBust() {
			gs.FlipCards()
			fmt.Print(gs.Print())
			fmt.Println("Dealer busts!")
			return ResultWin
		}
		if dealerStand {
			break
		}

	}
	gs.FlipCards()
	fmt.Print(gs.Print())
	fmt.Printf("Score: %v\n", gs.Player.Hand.Score)

	return gs.CompareHands()
}

// todo: remove the print by returning a string and deciding what to do with it in the call
func (gs *Gamestate) CompareHands() Result {
	fmt.Printf("%v - %v\n", gs.Player.Hand.Score, gs.Dealer.Hand.Score)
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
