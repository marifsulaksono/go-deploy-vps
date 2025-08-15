package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

func ConnectRedis(ctx context.Context) *redis.Client {
	// Ambil host & port dari env (kalau kosong, pakai default)
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "127.0.0.1"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	addr := fmt.Sprintf("%s:%s", host, port)

	rdb := redis.NewClient(&redis.Options{
		Addr:        addr,
		Password:    os.Getenv("REDIS_PASSWORD"), // kalau Redis pakai password
		DB:          0,
		DialTimeout: 2 * time.Second,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Gagal konek ke Redis di %s: %v", addr, err)
	}

	fmt.Printf("Berhasil konek ke Redis di %s → %s\n", addr, pong)
	return rdb
}
