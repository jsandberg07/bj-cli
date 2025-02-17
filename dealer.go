package main

type Dealer struct {
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
