package ping

import (
	"io"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
)

func RunCmd(args []string, con io.ReadWriter) []byte {
	var buf []byte

	if len(args) > 1 {
		return coder.Encode(config.ErrWrongNumberOfArguments("ping"), false)
	}
	
	if len(args) == 1 {
		buf = coder.Encode(args[0], true)
	} else {
		buf = coder.Encode("pong", true)
	}

	return buf
}