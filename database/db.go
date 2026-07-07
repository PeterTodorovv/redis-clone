package database

import (
	"fmt"
	"redis/request"
)

type Value interface {
	GetType() string
}

type Database struct {
	data map[string]Value
}

func CreateDatabase() *Database {
	db := Database{data: make(map[string]Value, 0)}
	return &db
}

type StringValue string

func (s StringValue) GetType() string {
	return "string"
}

const (
	ok        = "OK"
	not_found = ""
)

func (db *Database) Set(r request.Request) (string, error) {

	key, value, err := getSetArguments(r.GetArgs())
	if err != nil {
		return "", err
	}

	_, exists := db.data[key]

	if exists {
		return "", fmt.Errorf("Key %s already exists", key)
	}

	db.data[key] = value

	return ok, nil
}

func (db *Database) Get(r request.Request) (StringValue, error) {
	value, ok := db.data[r.GetArgs()[0]]
	if ok == false {
		return "", nil
	}

	val, err := value.(StringValue)

	if err == false {
		return "", nil
	}

	return val, nil
}

func getSetArguments(args []string) (string, StringValue, error) {
	if len(args) != 2 {
		return "", "", fmt.Errorf("Missing key or value parameters")
	}

	return args[0], StringValue(args[1]), nil
}
