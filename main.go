package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/patrickap/runr/m/v2/cmd"
	"github.com/patrickap/runr/m/v2/internal/lock"
)

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		lock.Unlock()
		os.Exit(1)
	}()

	err := cmd.Execute()
	if err != nil {
		lock.Unlock()
		os.Exit(1)
	}
}
