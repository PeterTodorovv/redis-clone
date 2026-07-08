package requesthandler

import (
	"redis/database"
	"redis/request"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestHandler(t *testing.T) {
	db := database.CreateDatabase()
	req := request.CreateRequest("PING", []string{})
	res := HandleRequest(*req, db)
	assert.Equal(t, res, "+PONG\r\n")

	req = request.CreateRequest("PIN", []string{})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "-ERR unknown command\r\n")

	req = request.CreateRequest("ping", []string{})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "+PONG\r\n")

	req = request.CreateRequest("set", []string{"name", "Sam"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "+OK\r\n")

	res = HandleRequest(*req, db)
	assert.Equal(t, res, "+OK\r\n")

	req = request.CreateRequest("get", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "$3\r\nSam\r\n")

	req = request.CreateRequest("get", []string{"age"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "$-1\r\n")

	req = request.CreateRequest("exists", []string{"age"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, ":0\r\n")

	req = request.CreateRequest("exists", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, ":1\r\n")

	req = request.CreateRequest("exists", []string{"name", "name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, ":2\r\n")

	req = request.CreateRequest("echo", []string{"test"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "$4\r\ntest\r\n")

	req = request.CreateRequest("del", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, ":1\r\n")

	req = request.CreateRequest("del", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, ":0\r\n")

	req = request.CreateRequest("get", []string{"name", "extra"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "-ERR wrong number of arguments for 'GET' command\r\n")

	req = request.CreateRequest("echo", []string{"a", "b"})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "-ERR wrong number of arguments for 'ECHO' command\r\n")

	req = request.CreateRequest("echo", []string{})
	res = HandleRequest(*req, db)
	assert.Equal(t, res, "-ERR wrong number of arguments for 'ECHO' command\r\n")
}
