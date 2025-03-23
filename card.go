package main

import "fmt"

type Suit int

const (
	SuitSpade Suit = iota
	SuitClub
	SuitHeart
	SuitDiamond
)

type Face int

const (
	FaceAce Face = iota
	FaceTwo
	FaceThree
	FaceFour
	FaceFive
	FaceSix
	FaceSeven
	FaceEight
	FaceNine
	FaceTen
	FaceJack
	FaceQueen
	FaceKing
)

type Visible int

const (
	VisibleFaceup Visible = iota
	VisibleFacedown
)

type Card struct {
	Suit
	Face
	Visible
}

func (c *Card) GetString() string {
	p := make([]byte, 2)
	switch c.Face {
	case FaceAce:
		p[0] = 'A'
	case FaceTwo:
		p[0] = '2'
	case FaceThree:
		p[0] = '3'
	case FaceFour:
		p[0] = '4'
	case FaceFive:
		p[0] = '5'
	case FaceSix:
		p[0] = '6'
	case FaceSeven:
		p[0] = '7'
	case FaceEight:
		p[0] = '8'
	case FaceNine:
		p[0] = '9'
	case FaceTen:
		p[0] = '0'
	case FaceJack:
		p[0] = 'J'
	case FaceQueen:
		p[0] = 'Q'
	case FaceKing:
		p[0] = 'K'
	default:
		p[0] = '?'
	}
	switch c.Suit {
	case SuitSpade:
		p[1] = 'S'
	case SuitClub:
		p[1] = 'C'
	case SuitHeart:
		p[1] = 'H'
	case SuitDiamond:
		p[1] = 'D'
	default:
	}
	return string(p)
}

func (c *Card) IsAce() bool {
	if c.Face == FaceAce {
		return true
	} else {
		return false
	}
}

func (c *Card) GetValue() int {
	switch c.Face {
	case FaceAce:
		return 11
	case FaceTwo:
		return 2
	case FaceThree:
		return 3
	case FaceFour:
		return 4
	case FaceFive:
		return 5
	case FaceSix:
		return 6
	case FaceSeven:
		return 7
	case FaceEight:
		return 8
	case FaceNine:
		return 9
	case FaceTen:
		fallthrough
	case FaceJack:
		fallthrough
	case FaceQueen:
		fallthrough
	case FaceKing:
		return 10
	default:
		fmt.Println("Default in GetValue")
		return 0
	}
}
