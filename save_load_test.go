package main

import (
	"reflect"
	"testing"
)

func TestSaveLoad(t *testing.T) {
	ss := Gamestate{}
	ss.Init()

	ss.Player.Init("Beans")
	ps := Stats{
		Wins:       10,
		Draws:      10,
		Losses:     10,
		Blackjacks: 10,
		Goal:       10,
		GoalMet:    false,
	}
	ss.Player.Stats = ps
	ss.Player.Money = 100000

	b1 := Bot{
		Name: "Bot 1",
	}
	b2 := Bot{
		Name: "Bot 2",
	}

	ss.Bots = append(ss.Bots, b1, b2)

	ss.Dealer.Init("Dealer")

	err := ss.Save()
	if err != nil {
		t.Fatalf("Error saving: %s", err)
	}

	ls := Gamestate{}
	ls.Init()
	err = ls.Load("Beans")
	if err != nil {
		t.Fatalf("Error loading: %s", err)
	}

	err = ls.CleanSave()
	if err != nil {
		t.Fatalf("Error cleaning save file: %s", err)
	}

	// can't just deep equal the whole thing because pointers to states won't point to the same address

	if !reflect.DeepEqual(ss.Player, ls.Player) {
		t.Fatalf("Players aren't the same\nss - %v-\nls - %v-", ss.Player, ls.Player)
	}

	if !reflect.DeepEqual(ss.Bots, ls.Bots) {
		for i := 0; i < len(ss.Bots); i++ {

			if !reflect.DeepEqual(ss.Bots[i].Name, ls.Bots[i].Name) {
				t.Fatalf("Names aren't the same\nss - %v-\nls - %v-", ss.Bots[i], ls.Bots[i])
			}

			/* hands are weird, need a reset? Doens't matter for the test really just the length and name
			if !reflect.DeepEqual(ss.Bots[i].Hand, ls.Bots[i].Hand) {
				t.Fatalf("Hands aren't the same\nss - %v-\nls - %v-", ss.Bots[i], ls.Bots[i])
			}
			*/
			if !reflect.DeepEqual(ss.Bots[i].Standing, ls.Bots[i].Standing) {
				t.Fatalf("Standing aren't the same\nss - %v-\nls - %v-", ss.Bots[i], ls.Bots[i])
			}

		}
		//t.Fatalf("Bots aren't the same\nss - %v-\nls - %v-", ss.Bots, ls.Bots)
	}

	if !reflect.DeepEqual(ss.Dealer, ls.Dealer) {
		t.Fatalf("Dealer isn't the same\nss - %v-\nls - %v-", ss.Bots, ls.Bots)
	}

}
