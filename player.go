package main

import (
	"fmt"
	"os"
)

type Player struct {
	Probability Probability
	Name        string
	Hand        Hand
	Money       int
	Bet         int
}

func (p *Player) TakeCard(c Card) {
	p.Hand.TakeCard(c)
}

func (p *Player) ResetHand() {
	p.Hand.Reset()
}

func (p *Player) IsBust() bool {
	return p.Hand.IsBust()
}

// TODO: handle input error gracefully
func (p *Player) GetPlayerChoice() Command {
	for {
		choice, err := getInput()
		if err != nil {
			fmt.Printf("Error getting input -- %s", err)
			os.Exit(1)
		}
		switch choice {
		case "stand":
			fallthrough
		case "s":
			fmt.Println("stand")
			return CommandStand
		case "hit":
			fallthrough
		case "h":
			fmt.Println("hit")
			return CommandHit
		default:
			fmt.Println("Unrecognized command")
		}
	}
}

func (p *Player) Init(name string) {
	p.Name = name
	p.Money = 200
	p.Bet = 0
	p.Hand.Init()
	p.Probability.Init(1)
}

func (p *Player) Reset(numDecks int) {
	p.Hand.Reset()
	p.Probability.Reset(numDecks)
}

// return lines that the game state will format
// name
// hand
func (p *Player) Print() []string {
	printName := p.Name
	printCards := p.Hand.GetCards()
	// if name is longer than cards, extend cards
	// if cards are longer than name, extend name
	// if equal, return early
	if len(printName) == len(printCards) {
		strings := []string{printName, printCards}
		return strings
	}

	if len(printName) > len(printCards) {
		printCards = centerPad(printCards, len(printName))
	} else {
		printName = centerPad(printName, len(printCards))
	}

	strings := []string{printName, printCards}
	return strings
}

func (p *Player) Win() {
	p.Money += 2 * p.Bet
	fmt.Println("You win!")
}

func (p *Player) Draw() {
	p.Money += p.Bet
	fmt.Println("Draw!")
}

// add check to see if you lose all your cash then exit
func (p *Player) Lose() {
	p.Money -= p.Bet
	fmt.Println("You lose!")
}
