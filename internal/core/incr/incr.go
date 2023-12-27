package incr

import (
	"fmt"
	"strconv"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func RunCmd(args []string) []byte {
	if len(args) != 1 {
		return coder.Encode(config.ErrWrongNumberOfArguments("incr"), false)
	}

	k := args[0]
	obj := storage.Get(k)
	if obj == nil {
		obj = storage.Add("0", -1, config.OBJ_TYPE_STRING, config.OBJ_ENCODING_INT)
		storage.Put(k, obj)
	}

	if !storage.AssertType(obj.TypeEncoding, config.OBJ_TYPE_STRING) {
		return coder.Encode(fmt.Errorf("ERR operation not permitted on type %d", storage.GetType(obj.TypeEncoding)), false)
	}

	if !storage.AssertEncoding(obj.TypeEncoding, config.OBJ_ENCODING_INT) {
		return coder.Encode(fmt.Errorf("ERR invalid encoding for key %s", k), false)
	}

	i, _ := strconv.ParseInt(obj.Value.(string), 10, 64)
	i++
	obj.Value = strconv.FormatInt(i, 10)
	
	return coder.Encode(i, false)
}