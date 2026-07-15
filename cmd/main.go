package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"redis/database"
	"redis/request"
	"redis/requestHandler"
)

func main() {
	listener, err := net.Listen("tcp", ":4200")
	db := database.CreateDatabase()

	if err != nil {
		log.Fatal(err)
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
	buffered := bufio.NewReader(connection)

	for {
		req, err := request.RequestFromReader(buffered)

		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println(err)
			connection.Write([]byte(err.Error()))
			return
		}

		response := requestHandler.HandleRequest(*req, db)
		connection.Write([]byte(response))
	}
}
