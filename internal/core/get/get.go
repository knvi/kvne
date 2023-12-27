package get

import (
	"io"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) []byte {
	if len(args) != 1 {
		return coder.Encode(config.ErrWrongNumberOfArguments("get"), false)
	}

	k := args[0]
	obj := storage.Get(k)
	if obj == nil {
		return config.RESP_NIL
	}

	if storage.HasExpired(obj) {
		storage.Del(k)
		return config.RESP_NIL
	}

	return coder.Encode(obj.Value, false)
}