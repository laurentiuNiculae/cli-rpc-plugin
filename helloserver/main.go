package main

import (
	"fmt"
	"log"
	"net"

	cli "github.com/laurentiuNiculae/go-rpc-plugins/plugins/clicommand"
	"google.golang.org/grpc"
)

const (
	port = 9001
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	cli.RegisterCLICommandServer(s, CLIServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
