package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	server := Server{}

	var response *httptest.ResponseRecorder
	for i := 0; i < 10; i++ {
		request, _ := http.NewRequest(http.MethodGet, "/pingpong", nil)
		response = httptest.NewRecorder()
		server.ServeHTTP(response, request)
	}

	got := response.Body.String()
	want := "pong 9"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
