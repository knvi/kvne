package main

import (
	"flag"
	"fmt"

	"github.com/knvi/kvne/internal/config"
	"github.com/knvi/kvne/internal/server"
)

func init() {
	flag.StringVar(&config.Host, "host", "0.0.0.0", "Host to listen on")
	flag.IntVar(&config.Port, "port", 6379, "Port to listen on")
	flag.Parse()
}

func main() {
	fmt.Println("Starting a kvne database...");
	server.RunAsyncTCPServer();
}