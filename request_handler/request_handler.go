package requesthandler

import "redis/request"

const (
	PING   = "PING"
	SET    = "SET"
	GET    = "GET"
	DEL    = "DEL"
	EXISTS = "EXISTS"
	ECHO   = "ECHO"
)

func handleRequest(r request.Request) {

}
