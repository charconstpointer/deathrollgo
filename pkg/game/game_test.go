package game

import (
	"log"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	g := NewGame(1000)
	p := NewPlayer(1)
	err := g.AddPlayer(p)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	err = g.AddPlayer(p)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	players := g.GetPlayers()
	if len(players) != 1 {
		t.Errorf("Expected %d player, found %d", 1, len(players))
	}
}

func TestRemovePlayer(t *testing.T) {
	g := NewGame(1000)
	p := NewPlayer(1)
	err := g.AddPlayer(p)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	err = g.RemovePlayer(p)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	players := g.GetPlayers()
	if len(players) != 0 {
		t.Errorf("Expected %d player, found %d", 0, len(players))
	}
}

func TestGetPlayers(t *testing.T) {
	g := NewGame(1000)
	p := NewPlayer(1)
	err := g.AddPlayer(p)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	players := g.GetPlayers()
	if len(players) != 1 {
		t.Errorf("Expected %d player, found %d", 1, len(players))
	}
}

func TestNextPlayer(t *testing.T) {
	g := NewGame(1000)
	p1 := NewPlayer(1)
	g.AddPlayer(p1)
	next, err := g.NextPlayer()
	if next.Id != p1.Id {
		t.Errorf("Wrong player, expected %d", p1.Id)
	}
	p2 := NewPlayer(2)
	err = g.AddPlayer(p2)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	next, err = g.NextPlayer()
	if next.Id != p1.Id {
		t.Errorf("Wrong order")
	}
	_, _, err = g.Roll()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	next, err = g.NextPlayer()
	if next.Id != p1.Id {
		t.Errorf("Wrong order")
	}
}

func TestRoll(t *testing.T) {
	g := NewGame(1000)
	p1 := NewPlayer(1)
	p2 := NewPlayer(2)
	err := g.AddPlayer(p1)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	err = g.AddPlayer(p2)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	id, roll, err := g.Roll()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	if id != p1.Id {
		t.Errorf("Wrong player, expected %d", p1.Id)
	}

	if roll < 0 {
		t.Errorf("Expected roll to be greater than 0, got %d", roll)
	}

}
