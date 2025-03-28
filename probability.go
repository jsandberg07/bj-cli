package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type Probability struct {
	// possibly better using the enum, but those dont have inherent values
	CardsLeft map[int]int
}

func (p *Probability) Init(numDecks int) {
	p.CardsLeft = map[int]int{}
	p.Reset(1)
}

func (p *Probability) Reset(numOfDecks int) {
	maps.Clear(p.CardsLeft)
	p.CardsLeft[11] = 4 * numOfDecks
	p.CardsLeft[2] = 4 * numOfDecks
	p.CardsLeft[3] = 4 * numOfDecks
	p.CardsLeft[4] = 4 * numOfDecks
	p.CardsLeft[5] = 4 * numOfDecks
	p.CardsLeft[6] = 4 * numOfDecks
	p.CardsLeft[7] = 4 * numOfDecks
	p.CardsLeft[8] = 4 * numOfDecks
	p.CardsLeft[9] = 4 * numOfDecks
	// because j, q, k, 0 all have a value of 10
	p.CardsLeft[10] = 4 * 4 * numOfDecks

}

// pass in current hand value, get percent to hit 21 and percent to bust
func (p *Probability) GetOdds(current int) (float64, float64) {
	toTarget := scoreTarget - current
	cardsTotal := 0
	cardsTarget := 0
	cardsBust := 0

	for key, value := range p.CardsLeft {
		if key == toTarget {
			cardsTarget += value
		}
		if key > toTarget {

			cardsBust += value
		}
		cardsTotal += value
	}
	// fmt.Println(p.CardsLeft)
	if cardsTotal == 0 {
		// no cards left. Shouldn't happen but this will keep it from crashing
		return 0.0, 0.0
	}
	targetOdds := float64(cardsTarget) / float64(cardsTotal)
	bustOdds := float64(cardsBust) / float64(cardsTotal)
	return targetOdds, bustOdds
}

// pass in value of card, not face
func (p *Probability) RemoveCard(value int) {
	p.CardsLeft[value]--
	if p.CardsLeft[value] < 0 {
		fmt.Println("Negative value found after removing card. Make sure to reset probability!")
	}
}
