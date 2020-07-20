package api

import (
	"context"
	"fmt"

	"github.com/charconstpointer/deathrollgo/pkg/game"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	game *game.Game
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error) {
	err := s.game.Start()
	if err != nil {
		return nil, err
	}
	return &StartGameResponse{}, nil
}

func (s *Server) CreateGame(c context.Context, r *CreateGameRequest) (*CreateGameResponse, error) {
	s.game = game.NewGame(int(r.Limit))
	log.Infof("Created new game with limit %d", r.Limit)
	return &CreateGameResponse{}, nil
}

func (s *Server) AddPlayer(c context.Context, r *AddPlayerRequest) (*AddPlayerResponse, error) {
	if s.game == nil {
		log.Errorf("You cannot add new player to the game that does not exists")
		return &AddPlayerResponse{
			Added: false,
		}, fmt.Errorf("You cannot add new player to the game that does not exists")
	}
	p := game.NewPlayer(r.UserId)

	err := s.game.AddPlayer(p)
	if err != nil {
		log.Errorf("Can't add player #%d, %v", r.UserId, err.Error())
		return &AddPlayerResponse{
			Added: false,
		}, fmt.Errorf("Can't add player #%d, %v", r.UserId, err.Error())
	}
	log.Infof("Adding player #%v", p)
	return &AddPlayerResponse{
		Added: true,
	}, nil
}

func (s *Server) Roll(ctx context.Context, r *RollRequest) (*RollResponse, error) {
	limit := s.game.GetLimit()

	roll, err := s.game.Roll(r.UserId)
	if err != nil {
		return nil, err
	}

	return &RollResponse{
		Value:  int32(roll),
		UserId: r.UserId,
		High:   int32(limit),
		Low:    0,
	}, nil
}

func (s *Server) GetNextPlayer(context.Context, *GetNextPlayerRequest) (*GetNextPlayerResponse, error) {
	if s.game == nil {
		log.Errorf("You cannot add new player to the game that does not exists")
		return nil, errors.Errorf("You cannot add new player to the game that does not exists")
	}
	next, err := s.game.NextPlayer()
	if err != nil {
		return &GetNextPlayerResponse{UserId: 0}, err
	}
	return &GetNextPlayerResponse{UserId: next.Id}, nil
}

func (s *Server) GetGameEvents(r *GetGameEventsRequest, stream GameService_GetGameEventsServer) error {
	if s.game == nil {
		return fmt.Errorf("Create game first in order to listen for game events")
	}

	for {
		select {
		case l := <-s.game.Losers:
			log.Infof("Event:Player %v lost", l)
			event := &GetGameEventsResponse{
				UserId: l.Id,
				Event:  GameEvent_PlayerLost,
			}
			if err := stream.Send(event); err != nil {
				return err
			}
		case w := <-s.game.Winner:
			event := &GetGameEventsResponse{
				UserId: w.Id,
				Event:  GameEvent_PlayerWon,
			}
			if err := stream.Send(event); err != nil {
				return err
			}
			return nil
		}
	}

}

func (s *Server) mustEmbedUnimplementedGameServiceServer() {

}
