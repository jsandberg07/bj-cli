package main

type Dealer struct {
	Name string
	Hand Hand
}

// dealer must draw to 16, stand on all 17s
const dealerTarget = 17

func (d *Dealer) MakeChoice() Command {
	if d.Hand.Score < dealerTarget {
		return CommandHit
	} else {
		return CommandStand
	}
}

func (d *Dealer) Init(name string) {
	d.Name = name
	d.Hand.Init()
}

func (d *Dealer) Reset() {
	d.Hand.Reset()
}

func (d *Dealer) Print() []string {
	printName := d.Name
	printCards := d.Hand.GetCards()
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

func (d *Dealer) TakeCard(c Card) {
	d.Hand.TakeCard(c)
}

func (d *Dealer) IsBust() bool {
	return d.Hand.IsBust()
}
