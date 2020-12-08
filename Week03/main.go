package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	g, _ := errgroup.WithContext(ctx)
	g.Go(func() error {
		return serveHttp(ctx.Done(), cancel)
	})

	g.Go(func() error {
		return handleSignal(ctx.Done(), cancel)
	})

	if err := g.Wait(); err != nil {
		fmt.Println("receive error", err)
	}

	time.Sleep(2 * time.Second) // 等待安全退出
}

func serveHttp(exitChan <-chan struct{}, cancelFunc context.CancelFunc) error {
	defer func() {
		log.Println("http server prepare exit")
		cancelFunc()
	}()

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	svc := &http.Server{
		Addr: ":8080",
	}

	go func() {
		select {
		case <-exitChan:
			ctx1, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			svc.Shutdown(ctx1)
		}
	}()

	err := svc.ListenAndServe()

	return fmt.Errorf("serve http failed: %w", err)
}

func handleSignal(exitChan <-chan struct{}, cancelFunc context.CancelFunc) error {
	defer func() {
		log.Println("handle signal prepare exit")
		cancelFunc()
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	select {
	case sign := <-quit:
		err := fmt.Errorf("handle signal: %d", sign)
		return err

	case <-exitChan:
		return nil
	}
}
