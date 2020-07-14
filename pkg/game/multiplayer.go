package game

type Multiplayer interface {
	AddPlayer(p *Player) error
	RemovePlayer(p *Player) error
}
