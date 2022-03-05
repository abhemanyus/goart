package goart_test

import (
	"net/http"
	"testing"

	"github.com/abhemanyus/goart"
	approvals "github.com/approvals/go-approval-tests"
)

func TestImageServer(t *testing.T) {
	database, _ := goart.CreateImageStore(fs)
	server := goart.CreateImageServer(database)
	getResponse := getServerResponse(server)
	t.Run("get image html", func(t *testing.T) {
		response := getResponse("/browser")
		assertStatusCode(t, response.Code, http.StatusOK)
		approvals.VerifyString(t, response.Body.String())
	})
	t.Run("get more page that available", func(t *testing.T) {
		response := getResponse("/browser?page=100")
		assertStatusCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStaticServer(t *testing.T) {
	database, _ := goart.CreateImageStore(fs)
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
