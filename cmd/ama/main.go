package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rodrigomolter/ama-backend/internal/api"
	"github.com/rodrigomolter/ama-backend/internal/store/pgstore"
)

func main() {
	env := os.Getenv("NODE_ENV")
	if env != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		os.Getenv("AMA_DATABASE_USER"),
		os.Getenv("AMA_DATABASE_PASSWORD"),
		os.Getenv("AMA_DATABASE_HOST"),
		os.Getenv("AMA_DATABASE_PORT"),
		os.Getenv("AMA_DATABASE_NAME"),
	))
	if err != nil {
		panic(err)
	}
	// Solves SQLSTATE 42P05 problem when using supabase in transaction mode as mention here https://github.com/jackc/pgx/issues/1847
	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
