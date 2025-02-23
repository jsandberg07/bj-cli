package main

import (
	"strings"
	"testing"
)

// write exactly what you want and then figure out how to make it happen
// auto resizing, return an int from the player for the width
// see if you can just insert []bytes in
func TestPrintingHands(t *testing.T) {
	c1 := Card{
		Suit: SuitSpade,
		Face: FaceAce,
	}
	c2 := Card{
		Suit: SuitDiamond,
		Face: FaceTwo,
	}
	c3 := Card{
		Suit: SuitHeart,
		Face: FaceThree,
	}
	c4 := Card{
		Suit: SuitClub,
		Face: FaceFour,
	}
	// AS 2D 3H 4C
	cards := []Card{c1, c2, c3, c4}
	names := []string{"Player", "BJ", "Beans Johnson"}

	results := [][]string{}

	// get all the outputs, then compare
	for i := 0; i < len(names); i++ {
		// create player
		p := Player{}
		// give name
		p.Init(names[i])
		// give 2 cards
		p.TakeCard(cards[0])
		p.TakeCard(cards[1])
		// get results
		results = append(results, p.Print())
		// another card
		p.TakeCard(cards[2])
		// compare
		results = append(results, p.Print())
		// another card
		p.TakeCard(cards[3])
		// compare
		results = append(results, p.Print())
	}

	// AS 2D 3H 4C
	expected := [][]string{}
	expected = append(expected, []string{
		"Player",
		"AS 2D "})
	expected = append(expected, []string{
		" Player ",
		"AS 2D 3H"})
	expected = append(expected, []string{
		"  Player   ",
		"AS 2D 3H 4C"})
	expected = append(expected, []string{
		" BJ  ",
		"AS 2D"})
	expected = append(expected, []string{
		"   BJ   ",
		"AS 2D 3H"})
	expected = append(expected, []string{
		"    BJ     ",
		"AS 2D 3H 4C"})
	expected = append(expected, []string{
		"Beans Johnson",
		"    AS 2D    "})
	expected = append(expected, []string{
		"Beans Johnson",
		"  AS 2D 3H   "})
	expected = append(expected, []string{
		"Beans Johnson",
		" AS 2D 3H 4C "})

	for i := 0; i < len(expected); i++ {
		for j := 0; j < len(expected[i]); j++ {
			if expected[i][j] != results[i][j] {
				t.Fatalf("Strings do not match %v-%v Compare:\nE: --%s--\nO: --%s--\n", i+1, j+1, expected[i][j], results[i][j])
			}
		}
	}

}

func TestPrintingPlayers(t *testing.T) {

	c1 := Card{
		Suit: SuitSpade,
		Face: FaceAce,
	}
	c2 := Card{
		Suit: SuitDiamond,
		Face: FaceTwo,
	}
	c3 := Card{
		Suit: SuitHeart,
		Face: FaceThree,
	}
	c4 := Card{
		Suit: SuitClub,
		Face: FaceFour,
	}
	// AS 2D 3H 4C
	cards := []Card{c1, c2, c3, c4}

	// new gs
	gs := Gamestate{}
	gs.Init()

	gs.P.Init("Player")
	gs.D.Init("Dealer")
	// give cards
	gs.P.TakeCard(cards[0])
	gs.P.TakeCard(cards[1])
	gs.D.TakeCard(cards[2])
	gs.D.TakeCard(cards[3])

	// better to go line by line
	output := gs.Print()
	splitOutput := strings.Split(output, "\n")
	expected := "|Player|Dealer|\n|AS 2D |3H 4C |\n"
	splitExpected := strings.Split(expected, "\n")

	if len(splitOutput) != len(splitExpected) {
		t.Fatalf("Different number of lines: E: %v, O: %v", len(splitExpected), len(splitOutput))
	}
	for i := 0; i < len(splitExpected); i++ {
		if splitExpected[i] != splitOutput[i] {
			t.Fatalf("Compare strings: \nE:--%s--\nO:--%s--", splitExpected[i], splitOutput[i])
		}
	}
}
