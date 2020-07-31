package main

import (
	"net"

	"github.com/charconstpointer/deathrollgo/pkg/game/api"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {

	log.SetFormatter(&log.JSONFormatter{})

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	s := api.NewServer()
	gs := grpc.NewServer()

	api.RegisterGameServiceServer(gs, s)
	gs.Serve(lis)
}
