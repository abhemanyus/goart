package goart_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhemanyus/goart"
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

func jsonToList(r io.Reader) (goart.ImageList, error) {
	var images goart.ImageList
	err := json.NewDecoder(r).Decode(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}
