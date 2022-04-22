package config

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Connection *grpc.ClientConn

func InitgRPC() (err error) {
	Connection, err = grpc.Dial(":8200", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	return nil
}
