package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/card-dungeon/server-card/db"
	cardpb "github.com/card-dungeon/server-card/protos/v1/card"
)

var (
	ADDRESS = flag.String("address", "localhost", "The server address")
	PORT    = flag.Int("port", 10503, "The server port")
)

type cardServer struct {
	cardpb.UnimplementedCardServer
}

func (s *cardServer) GetCharacterCard(ctx context.Context, req *cardpb.CardId) (*cardpb.CharacterCard, error) {
	cardId := req.CardId
	for _, c := range db.CharCardList {
		if c.CardId != cardId {
			continue
		}
		return c, nil
	}
	return &cardpb.CharacterCard{}, nil
}

func (s *cardServer) GetCharacterCardList(ctx context.Context, req *cardpb.GetList) (*cardpb.CharacterCardList, error) {
	return &cardpb.CharacterCardList{CharacterCards: db.CharCardList}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(*ADDRESS+":%d", *PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cardpb.RegisterCardServer(grpcServer, &cardServer{})
	reflection.Register(grpcServer)
	log.Printf("start gRPC server on %d port", *PORT)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
