package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	// events
	var signals = make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(2)

	server.RunAsyncTCPServer();
	server.WaitForSignal(&wg, signals)
}