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
	"github.com/woxQAQ/upload-server/pkg/types"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	cfg := &types.AppConfig{}
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
