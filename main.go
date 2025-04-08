package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"github.com/woxQAQ/upload-server/pkg/config"
)

func main() {
	cfg := &config.AppConfig{}
	viper.AddConfigPath(".")
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.ReadInConfig()
	err := viper.Unmarshal(cfg)
	if err != nil {
		panic(err)
	}
	db := config.InitDb(cfg)
	g := config.InitApp(db, cfg)

	srv := http.Server{
		Addr:    ":8080",
		Handler: g,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
