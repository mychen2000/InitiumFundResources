package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/initiumfund/artis-go/config"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "./config/config.yaml", "path to our configuration yaml file")
	flag.Parse()

	cfg := config.LoadConfig(configPath)

	log, err := SetupLogger(cfg)
	if err != nil {
		panic("Error while opening config files: " + err.Error())
	}
	defer log.Sync()

	db, err := SetupDataBase(cfg, log)
	if err != nil {
		log.Fatal("Failed to Setup Database" + err.Error())
	}

	// 用于调用 alpaca API 的客户端，会被当作 parameter 传进所有的 controllers
	alpacaClient := config.AlpacaClient(cfg)

	r, err := SetupRouter(cfg, db, alpacaClient, log)
	if err != nil {
		fmt.Printf("error while setting up router %s", err)
	}



	// Listen and serve on 0.0.0.0:8080
	gracefulServe(":8080", r, log)
}

// Allows HTTP Server to exit gracefully

func gracefulServe(addr string, router http.Handler, log *zap.SugaredLogger) {
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	fmt.Printf("server running on %s\n", addr)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Warn("Shutting down server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Warn("server shutdown error:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Info("5 seconds timed out")
	}
	log.Info("server exited")
}
