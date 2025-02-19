package main

import (
	"fmt"
	"os"
)

type Player struct {
	Name string
	Hand Hand
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
		fmt.Println("Default in MakePlayerChoice")
		return CommandStand
	}
}

func (p *Player) Init(name string) {
	p.Name = name
	p.Hand.Init()
}

func (p *Player) Reset() {
	p.Hand.Reset()
}

func (p *Player) GetScore() int {
	total := 0
	for _, c := range p.Hand.Cards {
		total += c.GetValue()
	}
	return total
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
