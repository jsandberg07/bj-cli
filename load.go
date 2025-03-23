package main

import (
	"encoding/json"
	"errors"
	"os"
)

type LoadData struct {
	P struct {
		Name  string `json:"Name"`
		Stats struct {
			Wins       int  `json:"Wins"`
			Draws      int  `json:"Draws"`
			Losses     int  `json:"Losses"`
			Blackjacks int  `json:"Blackjacks"`
			Goal       int  `json:"Goal"`
			GoalMet    bool `json:"GoalMet"`
		} `json:"Stats"`
		Money int `json:"Money"`
	} `json:"P"`
	B struct {
		NumBots  int      `json:"NumBots"`
		BotNames []string `json:"BotNames"`
	} `json:"B"`
	GS struct {
		NumDecks   int    `json:"NumDecks"`
		DealerName string `json:"DealerName"`
	} `json:"GS"`
}

func (gs *Gamestate) Load(name string) error {
	filename := name + ".json"
	// don't need to close if using ReadFile
	dat, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	ld := LoadData{}
	err = json.Unmarshal(dat, &ld)
	if err != nil {
		return err
	}

	if len(ld.B.BotNames) != ld.B.NumBots {
		return errors.New("number of bots inconsistent")
	}

	tp := Player{}
	tp.Init(ld.P.Name)
	tp.Stats = Stats(ld.P.Stats)
	tp.Money = ld.P.Money
	gs.Player = tp

	for i := 0; i < ld.B.NumBots; i++ {
		tb := Bot{}
		tb.Init(ld.B.BotNames[i])
		gs.Bots = append(gs.Bots, tb)
	}

	td := Dealer{}
	td.Init(ld.GS.DealerName)
	gs.Dealer = td
	gs.NumDecks = ld.GS.NumDecks

	return nil

}

func (gs *Gamestate) CleanSave() error {
	filename := gs.Player.Name + ".json"
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}
