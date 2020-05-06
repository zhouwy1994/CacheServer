package http

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type cacheHandler struct {
	*Server
}

func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := r.Method
	if m == http.MethodPut {
		value,_ := ioutil.ReadAll(r.Body)
		if len(value) != 0 {
			if h.Set(key, value) != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}

	if m == http.MethodGet {
		value, err := h.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(value) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(value)
		return
	}

	if m == http.MethodDelete {
		if h.Del(key) != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{s}
}