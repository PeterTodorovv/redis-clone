package requesthandler

import (
	"fmt"
	"redis/database"
	"redis/request"
	"strings"
)

const (
	cmdPing   = "PING"
	cmdSet    = "SET"
	cmdGet    = "GET"
	cmdDel    = "DEL"
	cmdExists = "EXISTS"
	cmdEcho   = "ECHO"
)

const (
	unknownCommand         = "ERR unknown command"
	wrongNumberOfArguments = "ERR wrong number of arguments for '%s' command"
)

func HandleRequest(r request.Request, db *database.Database) string {
	switch strings.ToUpper(r.GetCommand()) {
	case cmdPing:
		return simpleFormat(ping())
	case cmdSet:
		key, value, err := getKeyValueArguments(r.GetArgs())
		if err != nil {
			return errorFormat(err.Error())
		}

		response, err := db.Set(key, value)
		if err != nil {
			return errorFormat(err.Error())
		}
		return simpleFormat(response)
	case cmdGet:
		key, err := getKeyArgument(r.GetArgs())
		if err != nil {
			return errorFormat(err.Error())
		}

		response, found, err := db.Get(key)
		if err != nil {
			return errorFormat(err.Error())
		}
		return bulkFormat(string(response), found)
	case cmdExists:
		if len(r.GetArgs()) == 0 {
			return errorFormat(fmt.Sprintf(wrongNumberOfArguments, cmdExists))
		}
		found := db.Exists(r.GetArgs())
		return numberFormat(found)
	case cmdDel:
		if len(r.GetArgs()) == 0 {
			return errorFormat(fmt.Sprintf(wrongNumberOfArguments, cmdDel))
		}
		deleted := db.Del(r.GetArgs())
		return numberFormat(deleted)
	case cmdEcho:
		if len(r.GetArgs()) != 1 {
			return errorFormat(fmt.Sprintf(wrongNumberOfArguments, cmdEcho))
		}
		return bulkFormat(r.GetArgs()[0], true)
	default:
		return errorFormat(unknownCommand)
	}
}

func ping() string {
	return "PONG"
}

func simpleFormat(r string) string {
	return fmt.Sprintf("+%s\r\n", r)
}

func bulkFormat(r string, found bool) string {
	if !found {
		return "$-1\r\n"
	}
	return fmt.Sprintf("$%d\r\n%s\r\n", len(r), r)
}

func numberFormat(count int) string {
	return fmt.Sprintf(":%d\r\n", count)
}
func errorFormat(err string) string {
	return fmt.Sprintf("-%s\r\n", err)
}

func getKeyValueArguments(args []string) (string, database.StringValue, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf(wrongNumberOfArguments, cmdSet)
	}

	return args[0], database.StringValue(args[1]), nil
}

func getKeyArgument(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf(wrongNumberOfArguments, cmdGet)
	}
	return args[0], nil
}
