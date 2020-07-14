package api

import (
	"context"

	"github.com/charconstpointer/deathrollgo/pkg/game"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	game *game.Game
}

func (s *Server) StartGame(c context.Context, r *StartGameRequest) (*StartGameResponse, error) {
	s.game = game.NewGame(int(r.Limit))
	log.Infof("Created new game with limit %d", r.Limit)
	return &StartGameResponse{}, nil
}
func (s *Server) AddPlayer(c context.Context, r *AddPlayerRequest) (*AddPlayerResponse, error) {
	if s.game == nil {
		log.Errorf("Game is not created")
		return &AddPlayerResponse{
			Added: false,
		}, nil
	}
	p := game.NewPlayer(r.UserId)

	err := s.game.AddPlayer(p)
	if err != nil {
		log.Errorf("Can't add player #%d, %v", r.UserId, err.Error())
		return &AddPlayerResponse{
			Added: false,
		}, nil
	}
	log.Infof("Adding player #%v", p)
	return &AddPlayerResponse{
		Added: true,
	}, nil
}
func (s *Server) Roll(context.Context, *RollRequest) (*RollResponse, error) {
	if s.game == nil {
		log.Errorf("Game is not created")
		return &RollResponse{
			Success: false,
		}, nil
	}
	uid, roll, err := s.game.Roll()
	limit := s.game.GetLimit()
	if err != nil {
		log.Errorf("%v", err.Error())
	}
	return &RollResponse{
		Value:   int32(roll),
		Success: true,
		UserId:  uid,
		High:    int32(limit),
		Low:     0,
		Error:   "",
	}, nil
}
func (s *Server) GetNextPlayer(context.Context, *GetNextPlayerRequest) (*GetNextPlayerResponse, error) {
	if s.game == nil {
		log.Errorf("Game is not created")
		return nil, errors.Errorf("Game is not created")
	}
	next, err := s.game.NextPlayer()
	if err != nil {
		return &GetNextPlayerResponse{UserId: 0}, err
	}
	return &GetNextPlayerResponse{UserId: next.Id}, nil
}

func (s *Server) mustEmbedUnimplementedGameServiceServer() {

}
