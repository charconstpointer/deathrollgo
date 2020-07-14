package main

import (
	"log"

	"github.com/charconstpointer/deathrollgo/pkg/game"
)

func main() {
	g := game.NewGame(1000)
	p1 := game.NewPlayer(1, "foo")
	p2 := game.NewPlayer(2, "bar")
	p3 := game.NewPlayer(3, "bar")
	g.AddPlayer(p1)
	g.AddPlayer(p2)
	g.AddPlayer(p3)

	for i := 0; i < 10; i++ {
		p, roll, err := g.Roll()
		if err != nil {
			return
		}
		if roll == 0 {
			log.Printf("Winner : %v", p)
			return
		}
		log.Printf("%v rolled %d", p, roll)
	}

}
