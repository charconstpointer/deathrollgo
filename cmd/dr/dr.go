package main

import (
	"github.com/charconstpointer/deathrollgo/pkg/game"
)

func main() {
	g := game.NewGame()
	p1 := game.NewPlayer(1, "foo")
	p2 := game.NewPlayer(2, "bar")
	g.AddPlayer(p1)
	g.AddPlayer(p2)

}
