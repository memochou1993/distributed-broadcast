package main

import "github.com/go-redis/redis/v8"

func NewRDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}