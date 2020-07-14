package main

import (
	"log"

	"github.com/charconstpointer/deathrollgo/pkg/game"
)

func main() {
	g := game.NewGame()
	p := game.NewPlayer(1, "foo")
	log.Printf("%v ,%v", g, p)
}
