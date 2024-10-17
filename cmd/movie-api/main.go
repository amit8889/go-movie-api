package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amit8889/go-movie-api/internal/config"
)

func main() {
	slog.Info("=======main function started=========")
	cfg := config.MustLoad()
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGALRM)
	server := http.Server{
		Addr: cfg.HttpServer.Addr,
	}
	slog.Info("storage initalize", slog.String("PORT", cfg.HttpServer.Addr), slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("====error in server setup==>", err)
			panic(err)
		}
	}()
	<-done
	slog.Info("Server stopped")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("Server shutdown successfully")
}
