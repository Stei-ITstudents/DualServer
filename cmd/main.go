package main

import (
	"github.com/DualServer/internal/server1"
	"github.com/DualServer/internal/server2"
	"github.com/sirupsen/logrus"
)

func main() {
	go func() {
		if err := server1.Start(); err != nil {
			logrus.WithError(err).Error("Server 1 failed to start")
		}
	}()
	go func() {
		if err := server2.Start(); err != nil {
			logrus.WithError(err).Error("Server 2 failed to start")
		}
	}()

	logrus.Info("Dual server is running")
	select {} // block forever
}
