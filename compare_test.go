package main

import "testing"

func TestCompare(t *testing.T) {
	gs := Gamestate{}
	gs.Init()

	type CompareTest struct {
		TestName       string
		PlayerScore    int
		DealerScore    int
		ExpectedResult Result
	}

	tests := []CompareTest{
		{
			TestName:       "Player wins",
			PlayerScore:    21,
			DealerScore:    18,
			ExpectedResult: ResultWin,
		},
		{
			TestName:       "Draw",
			PlayerScore:    19,
			DealerScore:    19,
			ExpectedResult: ResultDraw,
		},
		{
			TestName:       "Dealer wins",
			PlayerScore:    18,
			DealerScore:    19,
			ExpectedResult: ResultLose,
		},
	}

	for i, test := range tests {
		gs.Player.Score = test.PlayerScore
		gs.Dealer.Score = test.DealerScore
		result := gs.CompareHands()
		if result != test.ExpectedResult {

			t.Fatalf("Test %v failed - Name - %s", i+1, test.TestName)
		}
	}
}
