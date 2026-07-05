package main

import (
	"fmt"
	"net"
	"redis/request"
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
		}

		_, err = request.RequestFromReader(connection)

		if err != nil {
			fmt.Println(err)
			connection.Write([]byte(err.Error()))
			continue
		}

		connection.Write([]byte("success"))
	}

}
