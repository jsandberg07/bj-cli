package main

import "strings"

type Command int

const (
	CommandHit Command = iota
	CommandStand
	CommandSurrender
	CommandBlackjack
)

type Hand struct {
	Cards   []Card
	Score   int
	NumAces int
}

func (h *Hand) Init() {
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

func (h *Hand) GetCards() string {
	var output string
	for _, c := range h.Cards {
		if c.Visible == VisibleFaceup {
			output += c.GetString() + " "
		} else {
			output += "XX" + " "
		}
	}
	return strings.TrimSpace(output)
}

func (h *Hand) HasBlackjack() bool {
	if h.Score == scoreTarget {
		return true
	} else {
		return false
	}
}
