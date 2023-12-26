package core

import (
	"fmt"
	"io"

	"github.com/knvi/kvne/internal/core/get"
	"github.com/knvi/kvne/internal/core/ping"
	"github.com/knvi/kvne/internal/core/set"
)

func Exec(cmd *Command, con io.ReadWriter) error {
	switch cmd.Name {
	case "ping":
		return ping.Run(cmd.Arguments, con)
	case "get":
		return get.RunCmd(cmd.Arguments, con)
	case "set":
		return set.RunCmd(cmd.Arguments, con)
	}

	return fmt.Errorf("unknown command: %s", cmd.Name)
}