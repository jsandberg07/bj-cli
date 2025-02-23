package main

type Bot struct {
	Name     string
	Hand     Hand
	Standing bool
}

// can make adjustable with composed logic later
func (b *Bot) MakeChoice() Command {
	if b.Hand.Score < dealerTarget {
		return CommandHit
	} else {
		return CommandStand
	}
}

func (b *Bot) Init(name string) {
	b.Name = name
	b.Standing = false
	b.Hand.Init()
}

func (b *Bot) Reset() {
	b.Standing = false
	b.Hand.Reset()
}

func (b *Bot) Print() []string {
	printName := b.Name
	printCards := b.Hand.GetCards()
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

func (b *Bot) TakeCard(c Card) {
	b.Hand.TakeCard(c)
}

func (b *Bot) IsBust() bool {
	return b.Hand.IsBust()
}
