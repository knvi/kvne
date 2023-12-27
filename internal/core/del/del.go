package del

import (
	"io"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string, c io.ReadWriter) []byte {
	delCount := 0

	for _, k := range args {
		if ok := storage.Del(k); ok {
			delCount++
		}
	}

	return coder.Encode(delCount, false)
}