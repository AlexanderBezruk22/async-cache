package main

import (
	"awesomeProject1/cachemanager"
	"fmt"
	"time"
)

func main() {
	cache := cachemanager.InitCache(5*time.Minute, 10*time.Minute)

	mapa := make(map[string]interface{})
	mapa["name"] = "Ivan"
	mapa["age"] = 18
	mapa["email"] = "testcache@g.co"

	cache.Set("consumer1", mapa, time.Minute)

	i, ok := cache.Get("consumer1")
	if !ok {
		fmt.Println("failed to get consumer1: not found")
	}
	fmt.Println(i)

	_, ok = cache.Get("consumer2")
	if !ok {
		fmt.Println("failed to get consumer2: not found")
	}

	ok = cache.Delete("consumer1")
	if ok {
		fmt.Println("Successfully deleted consumer1")
	}

	ok = cache.Delete("consumer2")
	if ok {
		fmt.Println("Successfully deleted consumer2")
	} else {
		fmt.Println("Failed to delete consumer2: not found")
	}
}
