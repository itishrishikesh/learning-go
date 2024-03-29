package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/itishrishikesh/learning-go/microservices/protos"
	"github.com/itishrishikesh/learning-go/microservices/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	cs := server.NewCurrency(log)

	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
