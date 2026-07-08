package request

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Request struct {
	command string
	args    []string
}

func (r Request) GetCommand() string {
	return r.command
}

func (r Request) GetArgs() []string {
	return r.args
}

const invalidPrefix = "Expected prefix %q, got %q"

func RequestFromReader(reader *bufio.Reader) (*Request, error) {
	count, err := readLength(reader, '*')
	if err != nil {
		return nil, err
	}

	command, err := readNext(reader)

	if err != nil {
		return nil, err
	}

	count--
	args := make([]string, count)
	for i := 0; i < count; i++ {
		arg, err := readNext(reader)

		if err != nil {
			return nil, err
		}

		args[i] = arg
	}

	return CreateRequest(command, args), nil
}

func CreateRequest(c string, a []string) *Request {
	return &Request{command: c, args: a}
}
func readLength(reader *bufio.Reader, prefix byte) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	line = strings.TrimSpace(line)
	if len(line) == 0 || line[0] != prefix {
		return 0, fmt.Errorf(invalidPrefix, prefix, line)
	}

	return strconv.Atoi(line[1:])
}

func readNext(reader *bufio.Reader) (string, error) {

	size, err := readLength(reader, '$')
	if err != nil {
		return "", err
	}

	buffer := make([]byte, size+2)
	if _, err := io.ReadFull(reader, buffer); err != nil {
		return "", err
	}
	return string(buffer[:size]), err
}
