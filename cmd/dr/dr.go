package main

import (
	"errors"
	"flag"
	"net"
	"strconv"
	"strings"

	"github.com/charconstpointer/deathrollgo/pkg/game/api"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	addr, err := parseAddr()
	if err != nil {
		addr = "0.0.0.0:8080"
	}

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}

	s := api.NewServer()
	gs := grpc.NewServer()

	api.RegisterGameServiceServer(gs, s)
	gs.Serve(lis)
}

func parseAddr() (string, error) {
	port := flag.Int("port", 8080, "port you want to listen on for incoming grpc requests")
	flag.Parse()

	if *port < 1024 || *port > 49151 {
		return "", errors.New("port out of range")
	}
	sb := strings.Builder{}
	sb.WriteString("0.0.0.0")
	sb.WriteString(":")
	sb.WriteString(strconv.Itoa(*port))
	return sb.String(), nil
}
