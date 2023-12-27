package core

import (
	"fmt"
	"io"

	"github.com/knvi/kvne/internal/core/bgrewriteaof"
	"github.com/knvi/kvne/internal/core/del"
	"github.com/knvi/kvne/internal/core/expire"
	"github.com/knvi/kvne/internal/core/get"
	"github.com/knvi/kvne/internal/core/ping"
	"github.com/knvi/kvne/internal/core/set"
	"github.com/knvi/kvne/internal/core/ttl"
)

func Exec(cmd *Command, con io.ReadWriter) error {
	var res []byte

	switch cmd.Name {
	case "ping":
		res = ping.RunCmd(cmd.Arguments, con)
	case "get":
		res = get.RunCmd(cmd.Arguments, con)
	case "set":
		res = set.RunCmd(cmd.Arguments, con)
	case "ttl":
		res = ttl.RunCmd(cmd.Arguments, con)
	case "del":
		res = del.RunCmd(cmd.Arguments, con)
	case "expire":
		res = expire.RunCmd(cmd.Arguments, con)
	case "bgrewriteaof":
		res = bgrewriteaof.RunCmd(cmd.Arguments)
	default:
		return fmt.Errorf("ERR unknown command '%s'", cmd.Name)
	}

	_, err := con.Write(res)
	return err
}