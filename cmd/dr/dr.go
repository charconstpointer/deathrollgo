package main

import (
	"log"
	"net"

	"github.com/charconstpointer/deathrollgo/pkg/game/api"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := api.Server{}
	gs := grpc.NewServer()

	api.RegisterGameServiceServer(gs, &s)
	gs.Serve(lis)
}
