package game

import "errors"

type queue []Player
type Game struct {
	players queue
}

type Multiplayer interface {
	AddPlayer(p *Player) error
	RemovePlayer(p *Player) error
}

func (g *Game) AddPlayer(p *Player) error {
	for _, pl := range g.players {
		if p.id == pl.id {
			return errors.New("Player is already in the game")
		}
	}
	g.players = append(g.players, *p)
	return nil
}

func (g *Game) RemovePlayer(p *Player) error {
	for i, pl := range g.players {
		if p.id == pl.id {
			g.players = append(g.players[:i], g.players[i+1:]...)
		}
	}
	return errors.New("Player not found")
}

func (g *Game) GetPlayers() []Player {
	return g.players
}

func NewGame() *Game {
	return &Game{}
}
