package requesthandler

import (
	"fmt"
	"redis/request"
	"strings"
)

const (
	PING   = "PING"
	SET    = "SET"
	GET    = "GET"
	DEL    = "DEL"
	EXISTS = "EXISTS"
	ECHO   = "ECHO"
)

const (
	UNKNOWN_COMMAND = "ERR unknown command"
)

func HandleRequest(r request.Request) (string, error) {
	switch strings.ToUpper(r.GetCommand()) {
	case PING:
		return simpleFormat(ping()), nil
	default:
		return errorFormat(UNKNOWN_COMMAND), nil
	}
}

func ping() string {
	return "PONG"
}

func simpleFormat(r string) string {
	return fmt.Sprintf("+%s\r\n", r)
}

func errorFormat(err string) string {
	return fmt.Sprintf("-%s\r\n", err)
}
