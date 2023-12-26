package core

import (
	"fmt"

	"github.com/knvi/kvne/internal/coder"
)

func ParseCommand(data []byte) (*Command, error) {
    value, err := coder.Decode(data)
	if err != nil {
		return nil, err
	}

	fmt.Println(value)

	array := value.([]interface{})
	if len(array) == 0 {
		return nil, fmt.Errorf("ERR empty command")
	}

	tokens := make([]string, len(array))
	for i, v := range array {
		tokens[i] = v.(string)
	}

	return &Command{tokens[0], tokens[1:]}, nil
}