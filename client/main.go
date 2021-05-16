package main

import (
	"fmt"
	"log"

	pb "github.com/cfabrica46/chat-grpc/chat"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:8080", grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	pb.NewChatServiceClient(conn)

	fmt.Println()

	fmt.Println()
}
