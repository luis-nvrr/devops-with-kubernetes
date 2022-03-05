package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	requestCount int
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/pingpong" {
		fmt.Fprintf(w, "pong %d", s.requestCount)
		s.requestCount++
	}
}
