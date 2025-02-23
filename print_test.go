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

	gs.Player.Init("Player")
	gs.Dealer.Init("Dealer")
	// give cards
	gs.Player.TakeCard(cards[0])
	gs.Player.TakeCard(cards[1])
	gs.Dealer.TakeCard(cards[2])
	gs.Dealer.TakeCard(cards[3])

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

func TestPrintingBots(t *testing.T) {
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
	c5 := Card{
		Suit: SuitSpade,
		Face: FaceFive,
	}
	c6 := Card{
		Suit: SuitDiamond,
		Face: FaceSix,
	}
	c7 := Card{
		Suit: SuitHeart,
		Face: FaceSeven,
	}
	c8 := Card{
		Suit: SuitClub,
		Face: FaceEight,
	}
	// AS 2D 3H 4C 5S 6D 7H 8C
	cards := []Card{c1, c2, c3, c4, c5, c6, c7, c8}

	gs := Gamestate{}
	gs.Init()
	gs.Player.Init("Player")
	gs.Dealer.Init("Dealer")
	gs.AddBots(2)

	gs.Player.TakeCard(cards[0])
	gs.Player.TakeCard(cards[1])
	gs.Dealer.TakeCard(cards[2])
	gs.Dealer.TakeCard(cards[3])
	gs.Bots[0].TakeCard(cards[4])
	gs.Bots[0].TakeCard(cards[5])
	gs.Bots[1].TakeCard(cards[6])
	gs.Bots[1].TakeCard(cards[7])

	output := gs.Print()
	splitOutput := strings.Split(output, "\n")
	expected := "|Player|Dealer|Bot 1|Bot 2|\n|AS 2D |3H 4C |5S 6D|7H 8C|\n"
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
