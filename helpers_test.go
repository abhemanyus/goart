package goart_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abhemanyus/goart"
)

func assertError(t testing.TB, got error, want error) {
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

func jsonToList(response *httptest.ResponseRecorder) (goart.ImageList, error) {
	var images goart.ImageList
	err := json.NewDecoder(response.Body).Decode(&images)
	return images, err
}
