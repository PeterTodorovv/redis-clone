package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPush(t *testing.T) {
	db := CreateDatabase()
	length, err := db.RPush("test", []string{"test"})
	assert.Len(t, db.data["test"], 1)
	assert.Equal(t, 1, length)
	assert.Nil(t, err)

	length, err = db.RPush("test", []string{"test", "test"})
	assert.Len(t, db.data["test"], 3)
	assert.Equal(t, 3, length)
	assert.Nil(t, err)

	db.data = map[string]Value{
		"key": StringValue("value"),
	}

	length, err = db.RPush("key", []string{"test"})
	assert.Equal(t, 0, length)
	assert.Error(t, err)

	db.data = map[string]Value{
		"key": ListValue{"test"},
	}

	db.RPush("key", []string{"test1"})

	assert.Equal(t, ListValue{"test", "test1"}, db.data["key"])
}

func TestLPush(t *testing.T) {
	db := CreateDatabase()
	length, err := db.RPush("test", []string{"test"})
	assert.Len(t, db.data["test"], 1)
	assert.Equal(t, 1, length)
	assert.Nil(t, err)

	length, err = db.RPush("test", []string{"test", "test"})
	assert.Len(t, db.data["test"], 3)
	assert.Equal(t, 3, length)
	assert.Nil(t, err)

	db.data = map[string]Value{
		"key": StringValue("value"),
	}

	length, err = db.RPush("key", []string{"test"})
	assert.Equal(t, 0, length)
	assert.Error(t, err)

	db.data = map[string]Value{
		"key": ListValue{"test"},
	}

	db.LPush("key", []string{"test1"})

	assert.Equal(t, ListValue{"test1", "test"}, db.data["key"])
}

func TestRPushX(t *testing.T) {
	db := CreateDatabase()
	length, err := db.RPush("test", []string{"test"})
	assert.Len(t, db.data["test"], 1)
	assert.Equal(t, 1, length)
	assert.Nil(t, err)

	length, err = db.RPush("test", []string{"test", "test"})
	assert.Len(t, db.data["test"], 3)
	assert.Equal(t, 3, length)
	assert.Nil(t, err)

	db.data = map[string]Value{
		"key": StringValue("value"),
	}

	length, err = db.RPush("key", []string{"test"})
	assert.Equal(t, 0, length)
	assert.Error(t, err)

	db.data = map[string]Value{
		"key": ListValue{"test"},
	}

	db.RPush("key", []string{"test1"})

	assert.Equal(t, ListValue{"test", "test1"}, db.data["key"])

	length, err = db.RPushX("key1", []string{"test1"})
	_, ok := db.data["key1"]
	assert.Equal(t, 0, length)
	assert.False(t, ok)

}

func TestLPushX(t *testing.T) {
	db := CreateDatabase()
	length, err := db.RPush("test", []string{"test"})
	assert.Len(t, db.data["test"], 1)
	assert.Equal(t, 1, length)
	assert.Nil(t, err)

	length, err = db.RPush("test", []string{"test", "test"})
	assert.Len(t, db.data["test"], 3)
	assert.Equal(t, 3, length)
	assert.Nil(t, err)

	db.data = map[string]Value{
		"key": StringValue("value"),
	}

	length, err = db.RPush("key", []string{"test"})
	assert.Equal(t, 0, length)
	assert.Error(t, err)

	db.data = map[string]Value{
		"key": ListValue{"test"},
	}

	db.LPush("key", []string{"test1"})

	assert.Equal(t, ListValue{"test1", "test"}, db.data["key"])
	length, err = db.LPushX("key1", []string{"test1"})
	_, ok := db.data["key1"]
	assert.Equal(t, 0, length)
	assert.False(t, ok)
}
