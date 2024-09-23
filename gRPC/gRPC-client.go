package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

var (
	client StorageClient
	conn   *grpc.ClientConn
	once   sync.Once
)

func Connect() StorageClient {
	once.Do(func() {
		var err error
		conn, err = grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed to connect to Python gRPC server at localhost:50051: %v", err)
		}
		client = NewStorageClient(conn)
	})
	return client
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}