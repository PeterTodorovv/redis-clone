package database

import "fmt"

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
	ok = "OK"
)

func (db *Database) Set(key string, value StringValue) (string, error) {
	db.data[key] = value

	return ok, nil
}

func (db *Database) Get(key string) (StringValue, bool, error) {
	value, ok := db.data[key]
	if !ok {
		return "", false, nil
	}

	val, ok := value.(StringValue)

	if !ok {
		return "", true, fmt.Errorf("Wrong datatype.")
	}

	return val, true, nil
}

func (db *Database) Exists(keys []string) int {
	found := 0

	for _, key := range keys {
		_, ok := db.data[key]

		if ok {
			found++
		}
	}

	return found
}
