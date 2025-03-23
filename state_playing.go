package main

import (
	"cmp"
	"fmt"
	"os"
)

func (gs *Gamestate) GetPlayingState() *State {
	s := State{
		Logic: PlayingLogic,
		// Print: PlayingPrint,
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
	case ResultBlackjack:
		gs.Player.Blackjack()
	case ResultSurrender:
		gs.Player.Surrender()
	default:
		fmt.Println("Default in playing logic!")
		os.Exit(1)
	}

	gs.Cleanup()

	gs.SetNextState(gs.getBettingState())

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
	fmt.Print(gs.PrintTable())
	// check if dealer or player has blackjack
	playerBJ := gs.Player.Hand.HasBlackjack()
	dealerBJ := gs.Player.Hand.HasBlackjack()
	if playerBJ && dealerBJ {
		return ResultDraw
	}
	if dealerBJ {
		return ResultLose
	}
	if playerBJ {
		return ResultBlackjack
	}

	// print probability
	toTarget, toBust := gs.Player.Probability.GetOdds(gs.Player.Hand.Score)
	fmt.Printf("To Target: %.2f%% - To Bust: %.2f%%\n", toTarget*100, toBust*100)
	// ask the player if he wants to hit or stand
	playerStand := false
	startingHand := true
	for {
		// player takes their turn
		fmt.Printf("Score: %v\n", gs.Player.Hand.Score)
		cmd := gs.Player.GetPlayerChoice()
		switch cmd {
		case CommandHit:
			startingHand = false
			gs.Player.TakeCard(gs.Deal(VisibleFaceup))
		case CommandStand:
			playerStand = true
		case CommandSurrender:
			if !startingHand {
				fmt.Println("Can only surrender on starting hand")
				continue
			}
			fmt.Println("Surrendering...")
			return ResultSurrender
		case CommandBlackjack:
			return ResultBlackjack
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

		fmt.Print(gs.PrintTable())

		if gs.Player.IsBust() {
			fmt.Println(gs.FlipCards())
			fmt.Print(gs.PrintTable())
			fmt.Printf("Bust! - %v\n", gs.Player.Hand.Score)
			return ResultLose
		}
		if playerStand {
			break
		}
		toTarget, toBust := gs.Player.Probability.GetOdds(gs.Player.Hand.Score)
		fmt.Printf("To Target: %.2f%% - To Bust: %.2f%%\n", toTarget*100, toBust*100)
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
			fmt.Println(gs.FlipCards())
			fmt.Print(gs.PrintTable())
			fmt.Println("Dealer busts!")
			return ResultWin
		}
		if dealerStand {
			break
		}

	}
	fmt.Println(gs.FlipCards())
	fmt.Print(gs.PrintTable())
	fmt.Printf("Score: %v\n", gs.Player.Hand.Score)

	return gs.CompareHands()
}

// todo: remove the print by returning a string and deciding what to do with it in the call
func (gs *Gamestate) CompareHands() Result {
	fmt.Printf("%v - %v\n", gs.Player.Hand.Score, gs.Dealer.Hand.Score)
	switch cmp.Compare(gs.Player.Hand.Score, gs.Dealer.Hand.Score) {
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
