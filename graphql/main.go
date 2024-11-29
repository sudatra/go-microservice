package main

import (
	"log"
	"net/http"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kelseyhightower/envconfig"
)

type AppConfig struct {
	AccountURL string `envConfig:"ACCOUNT_SERVICE_URL"`
	CatalogURL string `envConfig:"CATALOG_SERVICE_URL"`
	OrderURL   string `envConfig:"ORDER_SERVICE_URL"`
}

func main() {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	s, err := NewGraphQLServer(cfg.AccountURL, cfg.CatalogURL, cfg.OrderURL);
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/graphql", handler.New(s.ToExecutableSchema()));
	http.Handle("/playground", playground.Handler("sudatra", "/graphql"));

	log.Fatal(http.ListenAndServe(":8080", nil));
}