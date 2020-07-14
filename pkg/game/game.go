package game

import "errors"

type queue []Player
type Game struct {
	players queue
	actions int
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

func (g *Game) NextPlayer() *Player {
	pc := len(g.players)
	next := g.actions % pc
	return &g.players[next]
}

func (g *Game) Roll() (uint64, int) {
	current := g.NextPlayer()
	g.actions++
	return current.id, 1
}

func (g *Game) GetPlayers() []Player {
	return g.players
}

func NewGame() *Game {
	return &Game{}
}
