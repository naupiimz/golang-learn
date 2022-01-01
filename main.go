package main

import (
	"fmt"
)

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "zelda"}

	gamer.AddGame("Yordle")
	gamer.AddGame("Thrall")
	gamer.AddGame("Puffcaps")

	for _, games := range gamer.Games {
		fmt.Println(games)
	}
}
