package game

import (
	"log"
	"testing"
)

func TestAddPlayer(t *testing.T) {
	g := NewGame()
	p := NewPlayer(1, "foo")
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
	g := NewGame()
	p := NewPlayer(1, "foo")
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
	g := NewGame()
	p := NewPlayer(1, "foo")
	err := g.AddPlayer(p)
	if err != nil {
		log.Printf("%s", err.Error())
	}
	players := g.GetPlayers()
	if len(players) != 0 {
		t.Errorf("Expected %d player, found %d", 1, len(players))
	}
}
