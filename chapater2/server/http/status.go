package http

import (
	"encoding/json"
	"net/http"
)

type statusHandler struct {
	*Server
}

func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		body, err := json.Marshal(h.GetStat())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(body)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
