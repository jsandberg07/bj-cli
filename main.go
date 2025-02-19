package main

import (
	"fmt"
)

// break it all lmao

func main() {
	fmt.Println("Here we go")
	gs := Gamestate{}
	gs.Init()
	gs.Run()
	fmt.Println("Exiting...")
}
