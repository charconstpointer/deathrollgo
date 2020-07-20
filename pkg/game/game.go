package game

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type queue []Player
type Game struct {
	players  queue
	actions  int
	limit    int
	finished bool
	Losers   chan Player
	Winner   chan Player
}

func NewGame(limit int) *Game {
	return &Game{
		actions:  0,
		finished: false,
		limit:    limit,
		Losers:   make(chan Player),
		Winner:   make(chan Player, 1),
	}
}

func (g *Game) AddPlayer(p *Player) error {
	for _, pl := range g.players {
		if p.Id == pl.Id {
			return fmt.Errorf("Player is already in the game")
		}
	}
	g.players = append(g.players, *p)
	return nil
}

func (g *Game) RemovePlayer(p *Player) error {
	for i, pl := range g.players {
		if p.Id == pl.Id {
			g.players = append(g.players[:i], g.players[i+1:]...)
		}
	}
	return errors.New("Player not found")
}

func (g *Game) GetLimit() int {
	return g.limit
}

func (g *Game) NextPlayer() (*Player, error) {
	pc := len(g.players)
	if pc == 0 {
		return nil, errors.New("No players in the game")
	}

	next := g.actions % pc
	fmt.Printf("actions %d, pc %d\n", g.actions, pc)
	return &g.players[next], nil

}

func (g *Game) Roll(userId uint64) (int, error) {
	if g.finished {
		return 0, errors.New("This game is finished, start a new one to continue")
	}

	current, err := g.NextPlayer()
	if current.Id != userId {
		return 0, fmt.Errorf("Player %d is out of order", userId)
	}

	if err != nil {
		log.Printf("%s", err.Error())
	}

	pc := len(g.players)
	if current == nil {
		return 0, errors.New("Not enough players")
	}

	roll := rand.Intn(g.limit)
	//final stage
	if pc == 2 {
		g.headsUp(current, roll)
		return roll, nil
	}
	//picks a loser after full round
	if g.actions > 0 && g.actions%pc == 0 && pc > 2 {
		loser := g.pickLoser()
		log.Printf("%v lost", loser)
		g.Losers <- loser
		g.RemovePlayer(&loser)
		g.actions = 0
	}
	g.actions++
	current.score = roll
	return roll, nil
}

func (g *Game) headsUp(p *Player, roll int) {
	if roll == 0 {
		g.Losers <- *p
		g.RemovePlayer(p)
		w := g.players[0]
		g.Winner <- w
		g.finished = true
	}
	g.actions++
	g.limit = roll
}

func (g *Game) pickLoser() Player {
	var loser Player
	for i, p := range g.players {
		if i == 0 || p.score < loser.score {
			loser = p
		}
	}
	return loser
}

func (g *Game) GetPlayers() []Player {
	return g.players
}
