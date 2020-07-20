package game

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

type GameState int

const (
	NotStarted GameState = iota + 1
	InProgress
	Finished
)

type queue []Player
type Game struct {
	players queue
	actions int
	limit   int
	state   GameState
	Losers  chan Player
	Winner  chan Player
}

func NewGame(limit int) *Game {
	return &Game{
		actions: 0,
		state:   NotStarted,
		limit:   limit,
		Losers:  make(chan Player),
		Winner:  make(chan Player, 1),
	}
}

//AddPlayer checks if player is already present in the game, if not adds them
func (g *Game) Start() error {
	if g.state != NotStarted {
		return errors.New("You cannot start game that is already in progress or finished")
	}
	g.state = InProgress
	return nil
}

//AddPlayer checks if player is already present in the game, if not adds them
func (g *Game) AddPlayer(p *Player) error {
	if g.state != NotStarted {
		return errors.New("Can't join a game in progress")
	}
	for _, pl := range g.players {
		if p.Id == pl.Id {
			return errors.New("Player is already in the game")
		}
	}
	g.players = append(g.players, *p)
	return nil
}

//RemovePlayer removes player and decreases action count by one which is
//required for NextPlayer to work correctly
func (g *Game) RemovePlayer(p *Player) error {
	for i, pl := range g.players {
		if p.Id == pl.Id {
			g.players = append(g.players[:i], g.players[i+1:]...)
			g.actions--
		}
	}
	return errors.New("Player not found")
}

//GetLimit returns current upper bound value for Roll functionality
func (g *Game) GetLimit() int {
	return g.limit
}

//NextPlayer returns player that is expected to make the next Roll
func (g *Game) NextPlayer() (*Player, error) {
	pc := len(g.players)
	if pc == 0 {
		return nil, errors.New("No players in the game")
	}

	next := g.actions % pc
	fmt.Printf("actions %d, pc %d\n", g.actions, pc)
	return &g.players[next], nil

}

//Roll generates random roll within given boundary
func (g *Game) Roll(userId uint64) (int, error) {
	if g.state == Finished {
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
		g.state = Finished
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

//Returns all players in the game
func (g *Game) GetPlayers() []Player {
	return g.players
}
