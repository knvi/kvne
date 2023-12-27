package aof

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/knvi/kvne/internal/coder"
	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/core/storage"
)

func DumpAllAOF() {
	f, err := os.OpenFile(config.AOF_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Printf("Error opening file: %s", err)
	}
	defer f.Close()

	for k, o := range storage.Storage {
		cmd := fmt.Sprintf("SET %s %s", k, o.Value)
		tkns := strings.Split(cmd, " ")
		f.Write(coder.Encode(tkns, false))
	}

	log.Println("AOF dump completed")
}