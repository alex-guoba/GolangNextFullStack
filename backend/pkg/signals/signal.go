package signals

import (
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		log.Error("signal received, stopping...")
		close(stop)
		<-c
		log.Error("signal received, exit...")
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
