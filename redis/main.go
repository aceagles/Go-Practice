package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.140:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Context used to carry deadlines and timeouts etc for distributed proceses
	ctx := context.Background()

	session := map[string]string{"name": "John", "surname": "Smith", "company": "Redis", "age": "29"}
	for k, v := range session {
		err := client.HSet(ctx, "user-session:123", k, v).Err()
		if err != nil {
			panic(err)
		}
	}

	userSession := client.HGetAll(ctx, "user-session:123").Val()
	fmt.Println(userSession)
}
