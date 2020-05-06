package cache

import "log"

func New(typ string) Cache {
	var c Cache
	if typ == "inMemory" {
		c = newInMemoryCache()
	}

	if typ == "rocksdb" {
		c = newRocksdbCache()
	}

	if c == nil {
		panic("unknown cache type:" + typ)
	}

	log.Println(typ, "ready to server")
	return c
}
