package main

import (
	"context"
	"table_link/internal/infrastructure/db"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	ctx := context.Background()
	redisClient := db.Redis()
	postgres := db.DB()
}
