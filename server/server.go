package server

import (
	"context"
	"log"

        "github.com/izaaklauer/gpotato/config"
	gpotatov1 "github.com/izaaklauer/gpotato/gen/proto/go/gpotato/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GpotatoServer struct {
	gpotatov1.UnimplementedGpotatoServiceServer

	config config.Gpotato
}

// NewGpotatoServer initializes a new server from config
func NewGpotatoServer(config config.Gpotato) (*GpotatoServer, error) {
	// Server-specific initialization, like DB clients, goes here.

	server := GpotatoServer{
		config: config,
	}

	return &server, nil
}

func (s * GpotatoServer) HelloWorld(
	ctx context.Context,
	req *gpotatov1.HelloWorldRequest,
) (*gpotatov1.HelloWorldResponse, error) {
	log.Printf("HelloWorld request with message %q", req.Message)

	resp := &gpotatov1.HelloWorldResponse{
		RequestMessage: req.Message,
		ConfigMessage:  s.config.HelloWorldMessage,
		Now:            timestamppb.Now(),
	}

	return resp, nil
}
