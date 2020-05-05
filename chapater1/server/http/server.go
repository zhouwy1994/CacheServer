package http

import (
	"github.com/zhouwy1994/CacheServer/chapater1/server/cache"
	"net/http"
)

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	http.Handle("/cache/", s.cacheHandler())
	http.Handle("/status/", s.statusHandler())
	http.ListenAndServe(":3567", nil)
}

func New(c cache.Cache) *Server {
	return &Server{c}
}