package main

import (
	"fmt"
	"io"
	"log"

	"github.com/cfabrica46/crud-grpc/client/requests"
	pb "github.com/cfabrica46/crud-grpc/crud"
	"google.golang.org/grpc"
)

func main() {

	log.SetFlags(log.Lshortfile)

	conn, err := grpc.Dial("localhost:8080", grpc.WithBlock(), grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewCrudServiceClient(conn)

	//----

	var election int
	var exit bool

	for !exit {

		fmt.Println()
		fmt.Println("WELCOME!")
		fmt.Println("What would you like to do?")
		fmt.Println()

		fmt.Println("1.Log In")
		fmt.Println("2.Register")
		fmt.Println("0.Exit")
		fmt.Println()

		fmt.Print("> ")
		fmt.Scan(&election)
		fmt.Println()

		switch election {

		case 0:

			exit = true

		case 1:

			var username, password string

			fmt.Printf("\nENTER YOUR DATA\n")
			fmt.Printf("Username: ")
			fmt.Scan(&username)
			fmt.Printf("Password: ")
			fmt.Scan(&password)

			id, err := requests.Login(client, username, password)
			if err != nil {
				log.Fatal(err)
			}

			for !exit {
				err = loopIntoProfile(client, id, &exit)
				if err != nil {
					log.Fatal(err)
				}
			}

		case 2:

			var username, password string

			fmt.Printf("\nENTER YOUR DATA\n")
			fmt.Printf("Username: ")
			fmt.Scan(&username)
			fmt.Printf("Password: ")
			fmt.Scan(&password)

			err = requests.Register(client, username, password)

			if err != nil {
				log.Fatal(err)
			}

			id, err := requests.Login(client, username, password)
			if err != nil {
				log.Fatal(err)
			}

			for !exit {
				err = loopIntoProfile(client, id, &exit)
				if err != nil {
					log.Fatal(err)
				}
			}

		default:

			fmt.Println("Invalid Option")

		}
	}

}

func loopIntoProfile(client pb.CrudServiceClient, id int32, exit *bool) (err error) {

	var election int

	fmt.Println()
	fmt.Println("What would you like to do?")
	fmt.Println()

	fmt.Println("1.Show Users")
	fmt.Println("2.Delete Count")
	fmt.Println("0.Exit")
	fmt.Println()

	fmt.Print("> ")
	fmt.Scan(&election)
	fmt.Println()

	switch election {
	case 0:

		*exit = true

	case 1:

		err = requests.GetUsers(client)
		if err != nil {
			if err == io.EOF {
				err = nil
			} else {
				return
			}
		}

	case 2:

		err = requests.Delete(client, id)
		if err != nil {
			return
		}
		*exit = true

	default:

		fmt.Println("Seleccione una opción válida")

	}
	return
}
