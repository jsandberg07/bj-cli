package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	// ensure that two decks created aren't the same (are suffled upon init)
	d1 := new(Deck)
	d1.Init(1)
	d2 := new(Deck)
	d2.Init(1)
	if reflect.DeepEqual(d1.Cards, d2.Cards) {
		t.Fatal("Shuffle test 1 failed: decks are identical")
	}

	// create decks out of multiple and shuffle them
	d3 := new(Deck)
	d3.Init(8)
	d4 := new(Deck)
	d4.Init(8)
	if reflect.DeepEqual(d3.Cards, d4.Cards) {
		t.Fatal("Shuffle test 8 failed: decks are identical")
	}
}
