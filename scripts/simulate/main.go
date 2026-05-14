package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	redisIP := os.Getenv("REDIS_IP")
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{Addr: redisIP + ":6379"})

	go func() {
		for {
			for i := range 10 {
				key := fmt.Sprintf("product:%d", i)
				rdb.Get(ctx, key)
			}
			time.Sleep(time.Second)
		}
	}()

	fmt.Println("Phase 1: normal queries only (30s)")
	time.Sleep(30 * time.Second)

	fmt.Println("Phase 2: adding penetration queries")
	go func() {
		for {
			for i := range 10 {
				key := fmt.Sprintf("notexist%d", i)
				rdb.Get(ctx, key)
			}
			time.Sleep(time.Second)
		}
	}()

	select {}
}
