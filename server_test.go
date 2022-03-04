package goart_test

import (
	"net/http"
	"testing"

	"github.com/abhemanyus/goart"
)

func TestImageServer(t *testing.T) {
	database, _ := goart.CreateImageStore(fs)
	server := goart.CreateImageServer(database)
	getResponse := getServerResponse(server)
	t.Run("list images", func(t *testing.T) {
		response := getResponse("/list")
		images, err := jsonToList(response)
		assertError(t, err, nil)
		assertLength(t, 4, len(images))
	})
	t.Run("list 1 image", func(t *testing.T) {
		response := getResponse("/list?limit=1&offset=0")
		images, err := jsonToList(response)
		assertError(t, err, nil)
		assertLength(t, 1, len(images))
	})
	t.Run("get image", func(t *testing.T) {
		response := getResponse("/static/two.png")
		if response.Code != http.StatusOK {
			t.Errorf("requested image, wanted status code %d, got %d", http.StatusOK, response.Code)
		}
	})
}
