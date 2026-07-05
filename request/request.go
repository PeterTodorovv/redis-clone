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

const invlaid_prefix = "Expected prefix %q, got %q"

func readLength(reader *bufio.Reader, prefix byte) (int, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	line = strings.TrimSpace(line)
	if len(line) == 0 || line[0] != prefix {
		return 0, fmt.Errorf(invlaid_prefix, prefix, line)
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

func RequestFromReader(reader io.Reader) (*Request, error) {
	buffered := bufio.NewReader(reader)

	count, err := readLength(buffered, '*')
	if err != nil {
		return nil, err
	}

	request := &Request{}
	command, err := readNext(buffered)

	if err != nil {
		return nil, err
	}

	request.command = command
	for i := 1; i < count; i++ {
		arg, err := readNext(buffered)

		if err != nil {
			return nil, err
		}

		request.args = append(request.args, arg)
	}

	return request, nil
}
