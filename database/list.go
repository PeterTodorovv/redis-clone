package database

import "fmt"

type ListValue []string

const (
	list = "list"
)

func (l ListValue) GetType() string {
	return list
}

func (db *Database) RPush(key string, values []string) (int, error) {
	if _, ok := db.data[key]; !ok {
		db.data[key] = make(ListValue, 0)
	} else if db.data[key].GetType() != list {
		return 0, fmt.Errorf(wrongDatatype)
	}

	data := db.data[key]
	typeData := data.(ListValue)
	for _, value := range values {
		typeData = append(typeData, value)
		db.data[key] = typeData
	}

	return len(typeData), nil
}

func (db *Database) LPush(key string, values []string) (int, error) {
	if _, ok := db.data[key]; !ok {
		db.data[key] = make(ListValue, 0)
	} else if db.data[key].GetType() != list {
		return 0, fmt.Errorf(wrongDatatype)
	}

	data := db.data[key]
	typeData := data.(ListValue)
	for _, value := range values {
		typeData = append([]string{value}, typeData...)
		db.data[key] = typeData
	}

	return len(typeData), nil
}
