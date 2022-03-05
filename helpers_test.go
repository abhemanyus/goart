package goart_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("want error %d, but got %d", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("want error %q, but got %q", want, got)
	}
}

func getServerResponse(server http.Handler) func(string) *httptest.ResponseRecorder {
	return func(path string) *httptest.ResponseRecorder {
		request := httptest.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		return response
	}
}

func assertLength(t testing.TB, want, got int) {
	t.Helper()
	if got != want {
		t.Errorf("got incorrect number of images. Wanted %d, got %d", want, got)
	}
}
