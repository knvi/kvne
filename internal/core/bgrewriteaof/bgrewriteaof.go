package bgrewriteaof

import (
	"github.com/knvi/kvne/internal/aof"
	"github.com/knvi/kvne/internal/config"
)

func RunCmd(args []string) []byte {
	aof.DumpAllAOF()
	return config.RESP_OK
}