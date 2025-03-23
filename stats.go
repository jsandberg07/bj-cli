package main

import "fmt"

type Stats struct {
	Wins       int
	Draws      int
	Losses     int
	Blackjacks int
	// TODO: add back in but also add it to saving+loading
	// AverageBet AverageBet
	Goal    int
	GoalMet bool
}

func (s *Stats) Init() {
	s.Wins = 0
	s.Draws = 0
	s.Losses = 0
	s.Blackjacks = 0
	s.Goal = 0
	s.GoalMet = false
}

func (s *Stats) CheckGoal(m int) bool {
	if s.GoalMet {
		return true
	}
	if m > s.Goal {
		s.GoalMet = true
		return true
	} else {
		return false
	}
}

func (s *Stats) SetGoal(g int) {
	if g == 0 {
		s.Goal = 0
		s.GoalMet = true
	} else {
		s.Goal = g
		s.GoalMet = false
	}
}

func (s *Stats) PrintStats(money int) string {
	p := fmt.Sprintf("Wins - %v Draws - %v Losses - %v Blackjacks - %v", s.Wins, s.Draws, s.Losses, s.Blackjacks)
	if !s.GoalMet {
		p += fmt.Sprintf("\n%v until goal is met", s.Goal-money)
	}
	return p
}
