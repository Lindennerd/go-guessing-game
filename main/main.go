package main

import (
	"fmt"
	"guessing-game/game"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 10.")

	g := game.NewGame()
	g.Start()
}
