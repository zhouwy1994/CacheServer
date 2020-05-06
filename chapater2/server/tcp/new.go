package tcp

import (
	"github.com/zhouwy1994/CacheServer/chapater2/server/cache"
	"net"
)

type Server struct {
	cache.Cache
}

func New(c cache.Cache) *Server {
	return &Server{c}
}

func (s *Server) Listen() {
	listener, err := net.Listen("tcp4", ":3568")
	if err != nil {
		panic(err)
	}

	for {
		conn,err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go s.process(conn)
	}
}

