package main

import (
	"github.com/zhouwy1994/CacheServer/chapater1/server/cache"
	"github.com/zhouwy1994/CacheServer/chapater1/server/http"
)

func main() {
	c := cache.New("inMemory")
	s := http.New(c)
	s.Listen()
}
