package database

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

func (db *Database) Del(keys []string) int {
	deleted := 0
	for _, key := range keys {
		if _, ok := db.data[key]; ok {
			delete(db.data, key)
			deleted++
		}
	}
	return deleted
}
