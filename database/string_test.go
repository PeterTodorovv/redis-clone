package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	db := CreateDatabase()
	db.Set("test", StringValue("test"))
	_, ok := db.data["test"]
	assert.True(t, ok)
	assert.Equal(t, StringValue("test"), db.data["test"])

	db.Set("test", StringValue("value"))
	assert.Equal(t, StringValue("value"), db.data["test"])

	db.Set("test2", StringValue("test2"))
	assert.Len(t, db.data, 2)

	db.data = map[string]Value{
		"test": ListValue{},
	}

	db.Set("test", "test")
	assert.Equal(t, StringValue("test"), db.data["test"])
}

func TestGet(t *testing.T) {
	db := CreateDatabase()
	db.data = map[string]Value{
		"test": StringValue("test1"),
		"name": StringValue(""),
		"list": ListValue{},
	}
	value, found, err := db.Get("test")
	assert.Equal(t, StringValue("test1"), value)
	assert.True(t, found)
	assert.Nil(t, err)

	value, found, err = db.Get("name")
	assert.Equal(t, StringValue(""), value)
	assert.True(t, found)
	assert.Nil(t, err)

	value, found, err = db.Get("list")
	assert.NotNil(t, err)
	assert.Error(t, err)
}
