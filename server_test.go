package goart_test

import (
	"net/http"
	"testing"

	"github.com/abhemanyus/goart"
)

func TestAPIServer(t *testing.T) {
	database, err := goart.CreateImageStore(fs)
	assertError(t, err, nil)
	server := goart.CreateImageServer(database)
	getResponse := getServerResponse(server)
	t.Run("get <10 images", func(t *testing.T) {
		response := getResponse("/list?limit=10&offset=0")
		images, err := jsonToList(response.Body)
		assertError(t, err, nil)
		assertLength(t, 4, len(images))
	})
	t.Run("get >10 images", func(t *testing.T) {
		response := getResponse("/list?limit=100&offset=10")
		images, err := jsonToList(response.Body)
		assertError(t, err, nil)
		assertLength(t, 0, len(images))
	})
}

func TestImageServer(t *testing.T) {
	database, err := goart.CreateImageStore(fs)
	assertError(t, err, nil)
	server := goart.CreateImageServer(database)
	getResponse := getServerResponse(server)
	t.Run("get image", func(t *testing.T) {
		response := getResponse("/image/one.png")
		assertStatusCode(t, response.Code, http.StatusOK)
	})
	t.Run("get unavailable image", func(t *testing.T) {
		response := getResponse("/image/twotwo.png")
		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStaticServer(t *testing.T) {
	database, err := goart.CreateImageStore(fs)
	assertError(t, err, nil)
	server := goart.CreateImageServer(database)
	getResponse := getServerResponse(server)
	files := []string{
		"/static/browser.js",
		"/static/browser.css",
	}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			response := getResponse(file)
			assertStatusCode(t, response.Code, http.StatusOK)
		})
	}

}
