package cleanup

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func CloseResource() <-chan struct{} {
	onlyOneSignalHandler := make(chan struct{})
	close(onlyOneSignalHandler) // panics when called twice

	shutdownHandler := make(chan os.Signal, 2)

	ctx, cancel := context.WithCancel(context.Background())
	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM}
	signal.Notify(shutdownHandler, shutdownSignals...)
	go func() {
		<-shutdownHandler
		cancel()
		<-shutdownHandler
		os.Exit(1) // second signal. Exit directly.
	}()

	return ctx.Done()
}
