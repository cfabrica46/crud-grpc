package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/cfabrica46/crud-grpc/crud"
	"github.com/cfabrica46/crud-grpc/server/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type crudServer struct {
	pb.UnimplementedCrudServiceServer
}

func main() {

	l, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	creds, err := credentials.NewServerTLSFromFile("service.pem", "service.key")

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCrudServiceServer(grpcServer, &crudServer{})

	fmt.Println("Listening on localhost:8080")

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err)
	}

}

func (s *crudServer) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {

	var userAux *pb.User

	row := database.DB.QueryRow("SELECT users.id,users.role FROM users WHERE users.username = $1 AND users.password = $2", user.Username, user.Password)
	err := row.Scan(&userAux.ID, &userAux.Role)
	if err != nil {
		return nil, err
	}

	return userAux, nil
}

func (s *crudServer) GetUsers(user *pb.User, stream pb.CrudService_GetUsersServer) error {

	var userAux *pb.User

	rows, err := database.DB.Query("SELECT users.id,users.username,users.role FROM users")
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&userAux.ID, &userAux.Username, &userAux.Role)
		if err != nil {
			return err
		}

		if err = stream.Send(userAux); err != nil {
			return err
		}
	}

	return err
}

func (s *crudServer) AddUser(ctx context.Context, user *pb.User) (*pb.Message, error) {

	stmt, err := database.DB.Prepare("INSERT INTO users(username,password,role) VALUES ($1,$2,$3)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.Username, user.Password, "member")
	if err != nil {
		return nil, err
	}

	return &pb.Message{Body: "Se ah registrado con éxito"}, nil

}

func (s *crudServer) DeleteUser(ctx context.Context, user *pb.User) (*pb.Message, error) {

	stmt, err := database.DB.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(user.ID)
	if err != nil {
		return nil, err
	}

	return &pb.Message{Body: "Su cuenta ah sido elimindada con éxito"}, nil

}
