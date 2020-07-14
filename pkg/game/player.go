package game

type Player struct {
	id   uint64
	name string
}

func NewPlayer(id uint64, name string) *Player {
	return &Player{
		id:   id,
		name: name,
	}
}
