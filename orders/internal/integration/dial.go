package integration

import (
	"google.golang.org/grpc"
)

func Connect(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return conn, nil
}
