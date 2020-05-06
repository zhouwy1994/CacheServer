package main

import (
	"flag"
	"github.com/zhouwy1994/CacheServer/chapater3/server/cache"
	"github.com/zhouwy1994/CacheServer/chapater3/server/tcp"
	// "github.com/zhouwy1994/CacheServer/chapater3/server/http"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	c := cache.New(*typ)
	// go http.New(c).Listen()
	tcp.New(c).Listen()
}
