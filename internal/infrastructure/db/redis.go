package db

import (
	"os"

	redis "gopkg.in/redis.v5"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_HOST"),
	})

	check := client.Ping()
	if err := check.Err(); err != nil {
		panic(err)
	}

	return client
}
