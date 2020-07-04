package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_helloServer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/Foo", nil)
	response := httptest.NewRecorder()

	helloServer(response, request)

	got := response.Body.String()
	want := "Hello, Foo!"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
