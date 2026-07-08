package main

import (
	"fmt"
	"net"
	"redis/database"
	"redis/request"
	requesthandler "redis/request_handler"
)

func main() {
	listener, err := net.Listen("tcp", ":4200")
	db := database.CreateDatabase()

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		connection, err := listener.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		handleConnection(connection, db)
	}
}

func handleConnection(connection net.Conn, db *database.Database) {

	defer connection.Close()
	for {
		req, err := request.RequestFromReader(connection)

		if err != nil {
			fmt.Println(err)
			connection.Write([]byte(err.Error()))
			return
		}

		response := requesthandler.HandleRequest(*req, db)
		connection.Write([]byte(response))
	}
}
