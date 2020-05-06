package cache

// #include "rocksdb/c.h"
// #cgo CFLAGS: -I../../../rocksdb/include
// #cgo LDFLAGS: -L../../../rocksdb -lrocksdb -lz -lpthread -lsnappy -lstdc++ -lm -dl -O3
import "C"

type rocksdbCache struct {
	db *C.rocksdb_t
	ro *C.rocksdb_readoptions_t
	wo *C.rocksdb_writeoptions_t
	e  *C.char
}
