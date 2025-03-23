package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// add a goal first
func Fart() {
	fmt.Println("Pffrrtt.....")
}

// things to save
// name
// number of bots
// number of decks
// goal
// money
// stats
// i dunno everything
// just give it all a .save() and .load()
// when you get back from your walk
// ahh it wasn't typing cause i didn't have the window selected

// create a .json file
// save a .json file
// load a json file
// start putting data into it
// figure out how to structure things once you take it out
// saving is figuring out how to put everything into a struct
// loading is unpacking it
// game state needs to create a struct based on the player
// dealer
// bots (kinda)
// deck (kinda)
// mostly player
// so focus on that

type PlayerData struct {
	Name  string
	Stats Stats
	Money int
}

type BotData struct {
	NumBots  int
	BotNames []string
}

type GSData struct {
	NumDecks   int
	DealerName string
}

type SaveData struct {
	P  PlayerData
	B  BotData
	GS GSData
}

func (gs *Gamestate) Save() error {
	pd := PlayerData{
		Name:  gs.Player.Name,
		Stats: gs.Player.Stats,
		Money: gs.Player.Money,
	}

	botNames := []string{}
	if len(gs.Bots) > 0 {
		for i := 0; i < len(gs.Bots); i++ {
			botNames = append(botNames, gs.Bots[i].Name)
		}
	}
	bd := BotData{
		NumBots:  len(gs.Bots),
		BotNames: botNames,
	}

	gd := GSData{
		NumDecks:   gs.NumDecks,
		DealerName: gs.Dealer.Name,
	}

	sd := SaveData{
		P:  pd,
		B:  bd,
		GS: gd,
	}

	jj, err := json.Marshal(sd)
	if err != nil {
		return err
	}

	filename := gs.Player.Name + ".json"
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = os.WriteFile(filename, jj, os.FileMode(0644))
	if err != nil {
		return err
	}

	return nil

}
