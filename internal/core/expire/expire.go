package expire

import (
	"io"
	"time"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) []byte {
	if len(args) < 2 {
		return coder.Encode(config.ErrWrongNumberOfArguments("expire"), false)
	}

	k := args[0]
	obj := storage.Get(k)
	if obj == nil {
		return config.RESP_TTL_KEY_NOT_EXIST
	}

	exp, isSet := storage.GetExpiration(obj)
	if !isSet {
		return config.RESP_TTL_NO_EXPIRE
	}

	remain := exp - int64(time.Now().UnixMilli())
	if remain <= 0 {
		storage.Del(k)
		return config.RESP_TTL_KEY_NOT_EXIST
	}
	return coder.Encode(int64(remain/1000), true)
}