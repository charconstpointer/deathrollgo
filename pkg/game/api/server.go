package api

import (
	"context"
)

type Server struct {
}

func (s *Server) StartGame(context.Context, *StartGameRequest) (*StartGameResponse, error) {
	return &StartGameResponse{}, nil
}
func (s *Server) AddPlayer(context.Context, *AddPlayerRequest) (*AddPlayerResponse, error) {
	return &AddPlayerResponse{}, nil
}
func (s *Server) Roll(context.Context, *RollRequest) (*RollResponse, error) {
	return &RollResponse{}, nil
}
func (s *Server) mustEmbedUnimplementedGameServiceServer() {

}
