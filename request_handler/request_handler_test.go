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
	res, err := HandleRequest(*req, *db)
	assert.Nil(t, err)
	assert.Equal(t, res, "+PONG\r\n")

	req = request.CreateRequest("PIN", []string{})
	res, err = HandleRequest(*req, *db)
	assert.Equal(t, res, "-ERR unknown command\r\n")

	req = request.CreateRequest("ping", []string{})
	res, err = HandleRequest(*req, *db)
	assert.Equal(t, res, "+PONG\r\n")
}
