package main

import (
	"github.com/Saimunyz/L0/internal/migrations"
	"github.com/Saimunyz/L0/internal/model"
	"github.com/Saimunyz/L0/internal/nats"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"

	"github.com/Saimunyz/L0/internal/api"
	"github.com/Saimunyz/L0/internal/config"
	"github.com/Saimunyz/L0/pkg/helpers"
)

func main() {
	ctx := context.Background()
	ch := make(chan model.Data, 10)

	// read config.yml
	cfg, err := config.ReadConfigYML("config.yml")
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed init configuration due to error: %v", err))
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.SslMode,
	)

	// create postgres conn
	db, err := helpers.NewPostgres(ctx, dsn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed init postgres due to error: %v", err))
	}
	defer db.Close()

	// making migrations
	err = migrations.MakeMigrations(dsn, cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed init migrations due to error: %v", err))
	}

	// create nats conn
	sc, err := stan.Connect(cfg.Nats.ClusterID, cfg.Nats.ClientID)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed init nats conn due to err: %v", err))
	}
	defer helpers.Closer(sc)

	// Create sub on nats
	sub, err := nats.NewSubscription(sc, cfg, ch)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed init nats subscription due to err: %v", err))
	}
	defer sub.Unsubscribe()
	defer close(ch)

	router := httprouter.New()
	handler := api.NewHandler(db, ch)

	handler.Register(router)

	addr := fmt.Sprintf("%s:%s", cfg.HttpServer.IP, cfg.HttpServer.Port)
	log.Printf("Server is listening on %s\n", addr)

	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal(fmt.Sprintf("cannot create tcp connection due to err: %v", err))
	}
}
