package requestHandler

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
	assert.Equal(t, "+PONG\r\n", res)

	req = request.CreateRequest("PIN", []string{})
	res = HandleRequest(*req, db)
	assert.Equal(t, "-ERR unknown command\r\n", res)

	req = request.CreateRequest("ping", []string{})
	res = HandleRequest(*req, db)
	assert.Equal(t, "+PONG\r\n", res)

	req = request.CreateRequest("set", []string{"name", "Sam"})
	res = HandleRequest(*req, db)
	assert.Equal(t, "+OK\r\n", res)

	res = HandleRequest(*req, db)
	assert.Equal(t, "+OK\r\n", res)

	req = request.CreateRequest("get", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, "$3\r\nSam\r\n", res)

	req = request.CreateRequest("get", []string{"age"})
	res = HandleRequest(*req, db)
	assert.Equal(t, "$-1\r\n", res)

	req = request.CreateRequest("exists", []string{"age"})
	res = HandleRequest(*req, db)
	assert.Equal(t, ":0\r\n", res)

	req = request.CreateRequest("exists", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, ":1\r\n", res)

	req = request.CreateRequest("exists", []string{"name", "name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, ":2\r\n", res)

	req = request.CreateRequest("echo", []string{"test"})
	res = HandleRequest(*req, db)
	assert.Equal(t, "$4\r\ntest\r\n", res)

	req = request.CreateRequest("del", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, ":1\r\n", res)

	req = request.CreateRequest("del", []string{"name"})
	res = HandleRequest(*req, db)
	assert.Equal(t, ":0\r\n", res)

	req = request.CreateRequest("get", []string{"name", "extra"})
	res = HandleRequest(*req, db)
	assert.Equal(t, "-ERR wrong number of arguments for 'GET' command\r\n", res)

	req = request.CreateRequest("echo", []string{"a", "b"})
	res = HandleRequest(*req, db)
	assert.Equal(t, "-ERR wrong number of arguments for 'ECHO' command\r\n", res)

	req = request.CreateRequest("echo", []string{})
	res = HandleRequest(*req, db)
	assert.Equal(t, "-ERR wrong number of arguments for 'ECHO' command\r\n", res)
}
