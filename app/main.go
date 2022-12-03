package main

import (
	"context"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"latihan/app/bootstrap"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error load .env file")
	}
	server := bootstrap.InitHttp()
	runServer(server)
	waitForShutdown(server)
}

func runServer(srv *netHttp.Server) {
	go func() {
		log.Infoln("server run at", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != netHttp.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func waitForShutdown(server *netHttp.Server) {
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Fatal("system shutdown")

	// The context is used to inform the server it has 2 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("backend-service forced to shutdown")
	}

	log.Fatal("backend-service exiting")
}
