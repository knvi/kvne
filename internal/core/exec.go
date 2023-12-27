package core

import (
	"fmt"
	"io"

	"github.com/knvi/kvne/internal/core/del"
	"github.com/knvi/kvne/internal/core/expire"
	"github.com/knvi/kvne/internal/core/get"
	"github.com/knvi/kvne/internal/core/ping"
	"github.com/knvi/kvne/internal/core/set"
	"github.com/knvi/kvne/internal/core/ttl"
)

func Exec(cmd *Command, con io.ReadWriter) error {
	switch cmd.Name {
	case "ping":
		return ping.RunCmd(cmd.Arguments, con)
	case "get":
		return get.RunCmd(cmd.Arguments, con)
	case "set":
		return set.RunCmd(cmd.Arguments, con)
	case "ttl":
		return ttl.RunCmd(cmd.Arguments, con)
	case "del":
		return del.RunCmd(cmd.Arguments, con)
	case "expire":
		return expire.RunCmd(cmd.Arguments, con)
	}

	return fmt.Errorf("unknown command: %s", cmd.Name)
}