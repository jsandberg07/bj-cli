package main

import "testing"

func TestProp(t *testing.T) {
	p := Probability{}
	p.Init(1)
	p.Reset(1)

	type ProbabilityTest struct {
		Name           string
		Cards          []Card
		NumDecks       int
		PlayerHand     int
		ExpectedTarget float64
		ExpectedBust   float64
	}

	t1 := ProbabilityTest{
		Name: "No 5s",
		Cards: []Card{
			{Suit: SuitSpade, Face: FaceFive},
			{Suit: SuitDiamond, Face: FaceFive},
			{Suit: SuitHeart, Face: FaceFive},
			{Suit: SuitClub, Face: FaceFive},
		},
		PlayerHand:     21 - 5,
		NumDecks:       1,
		ExpectedTarget: 0,
		ExpectedBust:   36.0 / 48.0,
	}

	t2 := ProbabilityTest{
		Name: "One of each",
		Cards: []Card{
			{Suit: SuitClub, Face: FaceTwo},
			{Suit: SuitClub, Face: FaceThree},
			{Suit: SuitClub, Face: FaceAce},
			{Suit: SuitClub, Face: FaceFour},
			{Suit: SuitClub, Face: FaceFive},
			{Suit: SuitClub, Face: FaceSix},
			{Suit: SuitClub, Face: FaceSeven},
			{Suit: SuitClub, Face: FaceEight},
			{Suit: SuitClub, Face: FaceNine},
			{Suit: SuitClub, Face: FaceTen},
			{Suit: SuitClub, Face: FaceJack},
			{Suit: SuitClub, Face: FaceQueen},
			{Suit: SuitClub, Face: FaceKing},
		},
		PlayerHand:     18,
		NumDecks:       1,
		ExpectedTarget: 3.0 / 39.0,
		ExpectedBust:   33.0 / 39.0,
	}

	t3 := ProbabilityTest{
		Name: "No Hearts, Many Decks",
		Cards: []Card{
			{Suit: SuitHeart, Face: FaceTwo},
			{Suit: SuitHeart, Face: FaceTwo},
			{Suit: SuitHeart, Face: FaceTwo},
			{Suit: SuitHeart, Face: FaceThree},
			{Suit: SuitHeart, Face: FaceThree},
			{Suit: SuitHeart, Face: FaceThree},
			{Suit: SuitHeart, Face: FaceAce},
			{Suit: SuitHeart, Face: FaceAce},
			{Suit: SuitHeart, Face: FaceAce},
			{Suit: SuitHeart, Face: FaceFour},
			{Suit: SuitHeart, Face: FaceFour},
			{Suit: SuitHeart, Face: FaceFour},
			{Suit: SuitHeart, Face: FaceFive},
			{Suit: SuitHeart, Face: FaceFive},
			{Suit: SuitHeart, Face: FaceFive},
			{Suit: SuitHeart, Face: FaceSix},
			{Suit: SuitHeart, Face: FaceSix},
			{Suit: SuitHeart, Face: FaceSix},
			{Suit: SuitHeart, Face: FaceSeven},
			{Suit: SuitHeart, Face: FaceSeven},
			{Suit: SuitHeart, Face: FaceSeven},
			{Suit: SuitHeart, Face: FaceEight},
			{Suit: SuitHeart, Face: FaceEight},
			{Suit: SuitHeart, Face: FaceEight},
			{Suit: SuitHeart, Face: FaceNine},
			{Suit: SuitHeart, Face: FaceNine},
			{Suit: SuitHeart, Face: FaceNine},
			{Suit: SuitHeart, Face: FaceTen},
			{Suit: SuitHeart, Face: FaceTen},
			{Suit: SuitHeart, Face: FaceTen},
			{Suit: SuitHeart, Face: FaceJack},
			{Suit: SuitHeart, Face: FaceJack},
			{Suit: SuitHeart, Face: FaceJack},
			{Suit: SuitHeart, Face: FaceQueen},
			{Suit: SuitHeart, Face: FaceQueen},
			{Suit: SuitHeart, Face: FaceQueen},
			{Suit: SuitHeart, Face: FaceKing},
			{Suit: SuitHeart, Face: FaceKing},
			{Suit: SuitHeart, Face: FaceKing},
		},
		PlayerHand:     11,
		NumDecks:       3,
		ExpectedTarget: 36.0 / 117.0,
		ExpectedBust:   9.0 / 117.0,
	}

	t4 := ProbabilityTest{
		Name:           "No Decks at all!",
		Cards:          []Card{},
		PlayerHand:     20,
		NumDecks:       0,
		ExpectedTarget: 0.0,
		ExpectedBust:   0.0,
	}

	pts := []ProbabilityTest{t1, t2, t3, t4}

	for i, pt := range pts {
		// reset
		p.Reset(pt.NumDecks)
		// remove cards from probability
		for _, c := range pt.Cards {
			p.RemoveCard(c.GetValue())
		}

		outputTarget, outputBust := p.GetOdds(pt.PlayerHand)
		if outputTarget != pt.ExpectedTarget {
			t.Fatalf("Test: %v Name: %s\nTo target fail -- E: %.3f O: %.3f", i+1, pt.Name, pt.ExpectedTarget, outputTarget)
		}
		if outputBust != pt.ExpectedBust {
			t.Fatalf("Test: %v Name: %s\nTo bust fail -- E: %.3f O: %.3f", i+1, pt.Name, pt.ExpectedBust, outputBust)
		}
	}

}
