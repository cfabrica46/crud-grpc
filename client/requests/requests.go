package requests

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/cfabrica46/crud-grpc/crud"
)

func Login(client pb.CrudServiceClient, username, password string) (id int32, err error) {

	user := &pb.User{Username: username, Password: password}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err = client.GetUser(ctx, user)
	if err != nil {
		return
	}

	fmt.Printf("Bienvenido %v %v ID: %v\n", user.Role, user.Username, user.ID)

	id = user.ID

	return
}

func Register(client pb.CrudServiceClient, username, password string) (err error) {

	user := &pb.User{Username: username, Password: password}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message, err := client.AddUser(ctx, user)
	if err != nil {
		return
	}

	fmt.Println(message.Body)

	return
}

func GetUsers(client pb.CrudServiceClient) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := client.GetUsers(ctx, &pb.User{})
	if err != nil {
		return
	}

	for {
		var user *pb.User

		user, err = stream.Recv()

		if err != nil {
			if err == io.EOF {
				break
			}
			return
		}

		fmt.Printf("ID: %d\t|\tUsername: %s\t|\tRole: %s\n", user.ID, user.Username, user.Role)
	}

	return
}

func Delete(client pb.CrudServiceClient, id int32) (err error) {

	user := &pb.User{ID: id}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message, err := client.DeleteUser(ctx, user)
	if err != nil {
		return
	}

	fmt.Println(message.Body)

	return
}
