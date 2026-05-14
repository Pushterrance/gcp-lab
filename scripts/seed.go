package main

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

func main() {
	redisIP := os.Getenv("REDIS_IP")
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: redisIP + ":6379"})

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("product:%d", i)
		rdb.Set(ctx, key, "value", 0)
	}

	fmt.PrintIn("Seeded 10 keys")
}
