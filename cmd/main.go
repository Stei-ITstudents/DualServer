package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/CristyNel/DualServer/internal/server1"
	"github.com/CristyNel/DualServer/internal/server2"
	"github.com/sirupsen/logrus"
)

func main() {
	// *Create a context to handle graceful shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// *Create a channel to listen for signals.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	messageChan := make(chan string)

	var wgr sync.WaitGroup

	wgr.Add(1)

	server1Func := func() {
		if err := server1.Start(ctx, messageChan); err != nil {
			logrus.WithError(err).Error("Server 1 failed to start")
		}
	}
	go server1Func()

	wgr.Add(1)

	server2Func := func() {
		if err := server2.Start(ctx, messageChan); err != nil {
			logrus.WithError(err).Error("Server 2 failed to start")
		}
	}
	go server2Func()

	logrus.Info("Dual server is running")

	// *Message handling goroutine.
	go func() {
		for msg := range messageChan {
			logrus.Info(msg)
		}
	}()

	// *Wait for shutdown signal.
	select { // block until...
	case <-sigChan:
		logrus.Info("Received shutdown signal, stopping servers...")
		cancel() // Cancel the context to stop servers
	}

	wgr.Wait()

	logrus.Info("All servers have stopped.")
}
