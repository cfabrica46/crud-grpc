package main

import (
	"log"
	"net"

	pb "github.com/cfabrica46/chat-grpc/chat"
	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
}

func main() {

	l, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	grpcServer := grpc.NewServer()

}
