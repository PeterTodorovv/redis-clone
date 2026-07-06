package main

import (
	"fmt"
	"net"
	"redis/request"
	requesthandler "redis/request_handler"
)

func main() {
	listener, err := net.Listen("tcp", ":4200")
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

		req, err := request.RequestFromReader(connection)

		if err != nil {
			fmt.Println(err)
			connection.Write([]byte(err.Error()))
			continue
		}

		response, err := requesthandler.HandleRequest(*req)

		if err != nil {
			fmt.Println(err)
			connection.Write([]byte(err.Error()))
			continue
		}

		connection.Write([]byte(response))
	}

}
