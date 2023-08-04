package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"github.com/theLemonday/k8s-todos/api/handler"
	"github.com/theLemonday/k8s-todos/config"
	"github.com/theLemonday/k8s-todos/data/todo"
	"github.com/theLemonday/k8s-todos/database"
)

func main() {
	config := config.Load()
	fmt.Printf("%+v\n", config)

	mongoInstance, err := database.NewConnection(&config.MongodbConfig)
	if err != nil {
		log.Fatal().Err(err)
	}
	defer func() {
		mongoInstance.CloseConnection()
	}()
	mongoInstance.CheckConnection()
	todosColl := mongoInstance.GetCollection(config.MongodbConfig.Collection)

	repo := todo.NewMongoDbRepository(todosColl)
	service := todo.NewService(repo)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello world!"))
	})
	r.Mount("/todos", handler.TodosResource{}.Routes(service))

	startServer(&http.Server{
		Addr:    fmt.Sprintf(":%d", config.BackendConfig.Port),
		Handler: r,
	})
}

func startServer(server *http.Server) {
	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal().Msg("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal().Err(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}
