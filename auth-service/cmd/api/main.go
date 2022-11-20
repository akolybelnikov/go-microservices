package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/akolybelnikov/go-microservices/auth-service/data/users"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"net/http"
	"os"
	"time"
)

const webPort = "80"

type Config struct {
	DB      *sql.DB
	Queries *users.Queries
}

func main() {
	log.Println("Starting auth service")

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	app := Config{
		DB: db,
	}

	app.Queries = users.New(app.DB)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func openDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
