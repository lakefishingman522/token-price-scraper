package main

import (
	"context"
	// "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	db "github.com/CascadiaFoundation/CascadiaTokenScrapper/db"

	"github.com/CascadiaFoundation/CascadiaTokenScrapper/handlers"

	log "github.com/sirupsen/logrus"
)

func main() {

	logger := log.NewEntry(log.StandardLogger())

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	// Initialize a database
	DB, err := db.Init()

	if err != nil {
		panic(err)
	}

	h := handlers.New(DB)

	// Create routers
	router := gin.Default()

	// Create a new endpoint /health
	router.GET("/health", h.HealthCheck)

	// Create a new endpoint /getStatistics
	router.GET("/getStatistics", h.Statistics)

	logger.Info("after /getStatistics")

	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Create a timeout context used to gracefully shutdown the server with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()

	log.Println("Server exiting")
}
