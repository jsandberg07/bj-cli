package main

import (
	"fmt"
	"os"
)

type Command int

const (
	CommandHit Command = iota
	CommandStand
)

type Hand struct {
	Name    string
	Cards   []Card
	Score   int
	NumAces int
}

func (h *Hand) Init(name string) {
	h.Name = name
	h.Cards = []Card{}
	h.Score = 0
	h.NumAces = 0
}

func (h *Hand) Reset() {
	h.Cards = []Card{}
	h.Score = 0
	h.NumAces = 0
}

// if ace, add to ace and +value
// anything else, just add value
func (h *Hand) TakeCard(c Card) {
	if c.IsAce() {
		h.NumAces++
	}
	h.Score += c.GetValue()
	h.Cards = append(h.Cards, c)

	if h.Score > scoreTarget && h.NumAces > 0 {
		h.NumAces--
		h.Score -= 10
	}
}

func (h *Hand) IsBust() bool {
	if h.Score <= scoreTarget {
		return false
	} else {
		return true
	}
}

func (h *Hand) PrintHand() {
	fmt.Println(h.Name)
	fmt.Printf("Score: %v\n", h.Score)
	for _, c := range h.Cards {
		c.Print()
		fmt.Print(" ")
	}
	fmt.Println()
}

// temp for player making choice
// TODO: handle this gracefully
func (h *Hand) GetPlayerChoice() Command {
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

// temp for dealer making choice
func (h *Hand) MakeDealerChoice() Command {
	if h.Score < dealerTarget {
		return CommandHit
	} else {
		return CommandStand
	}
}
