package requesthandler

import (
	"fmt"
	"redis/database"
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

func HandleRequest(r request.Request, db database.Database) string {
	switch strings.ToUpper(r.GetCommand()) {
	case PING:
		return simpleFormat(ping())
	case SET:
		key, value, err := getSetArguments(r.GetArgs())
		if err != nil {
			return errorFormat(err.Error())
		}

		response, err := db.Set(key, value)
		if err != nil {
			return errorFormat(err.Error())
		}
		return simpleFormat(response)
	case GET:
		key, err := getGetKey(r.GetArgs())
		if err != nil {
			return errorFormat(err.Error())
		}

		response, err := db.Get(key)
		if err != nil {
			return errorFormat(err.Error())
		}
		return lenFormat(response)
	default:
		return errorFormat(UNKNOWN_COMMAND)
	}
}

func ping() string {
	return "PONG"
}

func simpleFormat(r string) string {
	return fmt.Sprintf("+%s\r\n", r)
}

func lenFormat(r database.StringValue) string {
	if r == "" {
		return "$-1\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(r), r)
}

func errorFormat(err string) string {
	return fmt.Sprintf("-%s\r\n", err)
}

func getSetArguments(args []string) (string, database.StringValue, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf("Missing key or value parameters")
	}

	return args[0], database.StringValue(args[1]), nil
}

func getGetKey(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("Missing get argument")
	}
	return args[0], nil
}
