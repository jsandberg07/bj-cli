package main

import (
	"fmt"
	"math/rand/v2"
)

type Deck struct {
	TopCard int
	Cards   []Card
}

func (d *Deck) Init() {
	d.TopCard = 0
	d.NewDeck()
	d.Shuffle()
}

func (d *Deck) Reset() {
	d.TopCard = 0
	d.Shuffle()
}

func (d *Deck) Print() {
	for _, c := range d.Cards {
		fmt.Println(c.GetString())
		fmt.Print(" ")
	}
}

func (d *Deck) Shuffle() {
	d.TopCard = 0
	t := make([]Card, len(d.Cards))
	for i, card := range rand.Perm(len(d.Cards)) {
		t[i] = d.Cards[card]
	}
	d.Cards = t
}

// create a single deck with 52 cards
func (d *Deck) NewDeck() {
	faces := []Face{FaceAce, FaceTwo, FaceThree, FaceFour, FaceFive, FaceSix, FaceSeven, FaceEight, FaceNine, FaceTen, FaceJack, FaceQueen, FaceKing}
	suits := []Suit{SuitSpade, SuitClub, SuitHeart, SuitDiamond}

	newDeck := make([]Card, 52)
	num := 0
	for i := 0; i < len(faces); i++ {
		for j := 0; j < len(suits); j++ {
			t := Card{
				Face: faces[i],
				Suit: suits[j],
			}
			newDeck[num] = t
			num++
		}
	}

	d.Cards = newDeck
}

// deal cards
func (d *Deck) Deal() Card {
	c := d.Cards[d.TopCard]
	d.TopCard++
	return c
}
