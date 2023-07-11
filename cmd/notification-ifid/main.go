package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rizalarfiyan/notification-ifid/internal"
)

func main() {
	port := 8080

	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	internal.NewRouter(route).All()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: route,
	}

	go func() {
		log.Printf("Server starting on port :%d\n", port)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	handleShutdown(server)
}

func handleShutdown(srv *http.Server) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
	return err
}
