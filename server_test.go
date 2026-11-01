package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETFoo(t *testing.T) {
	t.Run("returns 'bar'", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/foo", nil)
		response := httptest.NewRecorder()

		BarServer(response, request)

		got := response.Body.String()
		want := "bar"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
