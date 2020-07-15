package api

import (
	"context"

	"github.com/charconstpointer/deathrollgo/pkg/game"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	game   *game.Game
	Events chan (string)
}

func NewServer() *Server {
	return &Server{
		game:   game.NewGame(1000),
		Events: make(chan string),
	}
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

func (s *Server) GetGameEvents(r *GetGameEventsRequest, stream GameService_GetGameEventsServer) error {
	log.Printf("Listening for events")
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
			log.Infof("Event:Player %v won", w)
			event := &GetGameEventsResponse{
				UserId: w.Id,
				Event:  GameEvent_PlayerWon,
			}
			if err := stream.Send(event); err != nil {
				return err
			}
		}
	}

}

func (s *Server) mustEmbedUnimplementedGameServiceServer() {

}
