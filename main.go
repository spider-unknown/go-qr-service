package main

import (
	"context"
	httpserver "document-qr/services"
	"fmt"
	"github.com/unistack-org/micro/v3/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)

	logger.DefaultLogger = logger.NewLogger(logger.WithLevel(logger.TraceLevel))

	go httpserver.StartDocumentQRService(ctx, errCh)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errCh <- fmt.Errorf("%s", <-ch)
	}()

	logger.Infof(ctx, "Service terminated %v", <-errCh)
}
