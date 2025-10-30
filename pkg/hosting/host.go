package hosting

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Host struct {
	Workers []Worker
	Addr    string
	server  *http.Server
}

func (h *Host) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	wg := &sync.WaitGroup{}

	for _, w := range h.Workers {
		wg.Add(1)
		worker := w
		go worker.Run(ctx, wg.Done)
	}

	<-ctx.Done()
	log.Println("host shutdown started")
	defer log.Println("host shutdown completed")
	wg.Wait()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := h.server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server was forced to shutdown")
	}
}
