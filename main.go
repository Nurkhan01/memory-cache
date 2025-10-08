package main

import (
	"fmt"
	"golang-ninja/memoryCache"
)

func main() {
	cache := memoryCache.NewMemoryCache[float64]()
	cache.Set("dota2 version", 1.2)
	cache.Set("cs2 version", 1.5)
	fmt.Println(cache.Get("dota2 version"))
	fmt.Println(cache.Get("cs2 version"))
	cache.Del("dota2 version")
	fmt.Println(cache.Get("dota2 version"))
}
