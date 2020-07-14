package game

type Player struct {
	Id    uint64
	score int
}

func NewPlayer(id uint64) *Player {
	return &Player{
		Id: id,
	}
}
