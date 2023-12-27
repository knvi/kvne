package events

import "github.com/knvi/kvne/internal/core/bgrewriteaof"

func Shutdown() {
	bgrewriteaof.RunCmd([]string{})
}