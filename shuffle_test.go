package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	// ensure that two decks created aren't the same
	d1 := new(Deck)
	d1.Init()
	d1.Shuffle()
	d2 := new(Deck)
	d2.Init()
	d2.Shuffle()
	if reflect.DeepEqual(d1.Cards, d2.Cards) {
		t.Fatal("Shuffle test failed: decks are identical")
	}
}
