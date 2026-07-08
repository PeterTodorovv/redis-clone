package database

import "fmt"

const (
	ok = "OK"
)

type StringValue string

func (s StringValue) GetType() string {
	return "string"
}

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
