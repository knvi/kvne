package ping

import (
	"errors"
	"io"

	"github.com/knvi/kvne/internal/coder"
)

func RunCmd(args []string, con io.ReadWriter) error {
	var buf []byte

	if len(args) > 1 {
		return errors.New("ERR wrong number of arguments for 'ping' command")
	}
	
	if len(args) == 1 {
		buf = coder.Encode(args[0], true)
	} else {
		buf = coder.Encode("pong", true)
	}

	_, err := con.Write(buf)

	return err
}