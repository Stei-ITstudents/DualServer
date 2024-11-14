package server2

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	Port2 = 8082
	Sec5  = 5 * time.Second
	Sec10 = 10 * time.Second
	Sec15 = 15 * time.Second
)

func Start(ctx context.Context, messageChan chan string) error {
	http.HandleFunc("/server2", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Server 2 is running")
	})

	logrus.Infof("Starting server on http://localhost:%d/server2 ...", Port2)

	// *Periodically send status messages
	go func() {
		for {
			select {
			case <-ctx.Done():
				logrus.Info("Server 2 shutting down")

				return
			case <-time.After(Sec10):
				messageChan <- "Server 2 is alive!"
			}
		}
	}()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", Port2),
		ReadTimeout:  Sec5,
		WriteTimeout: Sec10,
		IdleTimeout:  Sec15,
	}

	// *Start the server and handle context cancellation
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	if err := server.ListenAndServe(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
