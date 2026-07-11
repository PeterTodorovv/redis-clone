package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	db := CreateDatabase()
	db.data = map[string]Value{
		"test":  StringValue("test"),
		"test1": StringValue("test1"),
	}

	found := db.Exists([]string{"test"})
	assert.Equal(t, 1, found)

	found = db.Exists([]string{"test2"})
	assert.Equal(t, 0, found)

	found = db.Exists([]string{"test", "test1"})
	assert.Equal(t, 2, found)

	found = db.Exists([]string{"test", "test1", "test1"})
	assert.Equal(t, 3, found)

	found = db.Exists([]string{"test", "test1", "test2"})
	assert.Equal(t, 2, found)
}

func TestDel(t *testing.T) {
	db := CreateDatabase()
	db.data = map[string]Value{
		"test":  StringValue("test"),
		"test1": StringValue("test1"),
	}

	deleted := db.Del([]string{"test"})
	assert.Equal(t, 1, deleted)

	deleted = db.Del([]string{"test"})
	assert.Equal(t, 0, deleted)

	deleted = db.Del([]string{"test1", "test1"})
	assert.Equal(t, 1, deleted)

	db.data = map[string]Value{
		"test":  StringValue("test"),
		"test1": StringValue("test1"),
	}

	deleted = db.Del([]string{"test", "test1"})
	assert.Equal(t, 2, deleted)
}
