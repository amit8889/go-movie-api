package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amit8889/go-movie-api/internal/config"
	"github.com/amit8889/go-movie-api/internal/http/router"
	"github.com/amit8889/go-movie-api/internal/storage/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	slog.Info("=======main function started=========")
	cfg := config.MustLoad()
	slog.Info("storage initalize", slog.String("PORT", cfg.HttpServer.Addr), slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
	fmt.Println(cfg)
	// db connection
	var storage *mongo.Database = mongodb.ConnectDb(cfg.MONGO_URL)
	router := router.MovieRouter(storage)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGALRM)
	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}
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
