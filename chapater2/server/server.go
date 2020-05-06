package main

import (
	"github.com/zhouwy1994/CacheServer/chapater2/server/cache"
	"github.com/zhouwy1994/CacheServer/chapater2/server/tcp"
)

func main() {
	c := cache.New("inmemory")
	s := tcp.New(c)
	s.Listen()
}
