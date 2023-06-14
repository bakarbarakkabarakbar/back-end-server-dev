package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	var ctx context.Context
	var err error
	var client *redis.Client
	var url = "localhost"
	var result string

	client, err = ConnectToRedis(&url, &ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = WriteCache(client, "hai", "syg")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result, err = RetrieveCache(client, "hai")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}

func WriteCache(client *redis.Client, key string, value string) error {
	var err error
	var ctx context.Context

	ctx = context.Background()

	err = client.Set(ctx, key, value, time.Hour*1).Err()
	if err != nil {
		return err
	}

	return nil
}

func RetrieveCache(client *redis.Client, key string) (string, error) {
	var err error
	var ctx context.Context
	var result string

	result, err = client.Get(ctx, "foo").Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func ConnectToRedis(url *string, ctx *context.Context) (*redis.Client, error) {
	var client *redis.Client
	var test string
	var err error
	client = redis.NewClient(&redis.Options{
		Addr:     *url + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	test, err = client.Ping(*ctx).Result()
	fmt.Println(test)
	if err != nil {
		return nil, err
	}
	return client, nil
}
