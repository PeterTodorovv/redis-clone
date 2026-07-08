package request

import (
	"bufio"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Reader struct {
	data   string
	isDone bool
}

func (reader *Reader) Read(data []byte) (int, error) {
	if reader.isDone {
		return 0, io.EOF
	}
	copy(data, []byte(reader.data))
	reader.isDone = true
	return len(reader.data), nil
}

func TestRequestParsing(t *testing.T) {
	reader := Reader{data: "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$5\r\nPeter\r\n"}
	buffered := bufio.NewReader(&reader)
	request, _ := RequestFromReader(buffered)

	assert.Equal(t, request.command, "SET")
	assert.Equal(t, request.args[0], "name")
	assert.Equal(t, request.args[1], "Peter")
	assert.Len(t, request.args, 2)

	reader = Reader{data: "*3\r\n$3\r\nSET\r\n$4\r\nname\r\n$5\r\nPeter\r\n$5\r\nPeter\r\n"}
	buffered = bufio.NewReader(&reader)
	request, _ = RequestFromReader(buffered)

	assert.Equal(t, request.command, "SET")
	assert.Equal(t, request.args[0], "name")
	assert.Equal(t, request.args[1], "Peter")
	assert.Len(t, request.args, 2)

	reader = Reader{data: "*1\r\n$4\r\nPING\r\n"}
	buffered = bufio.NewReader(&reader)
	request, _ = RequestFromReader(buffered)

	assert.Equal(t, request.command, "PING")
	assert.Len(t, request.args, 0)

}
