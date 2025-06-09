package db

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Init() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, continuing...")
    }

    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
        os.Getenv("SSL_MODE"),
    )

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    pool, err := pgxpool.New(ctx, dsn)
    if err != nil {
        log.Fatalf("Unable to connect to database: %v", err)
    }

    if err = pool.Ping(ctx); err != nil {
        log.Fatalf("Unable to ping database: %v", err)
    }

    DB = pool
    log.Println("Connected to PostgreSQL")
}
