package goart

import (
	"embed"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type ImageServer struct {
	http.Handler
}

var (
	//go:embed "static/*"
	staticFS embed.FS
)

func CreateImageServer(database *ImageDatabase) *ImageServer {
	server := &ImageServer{}
	router := createRouter(database)
	server.Handler = router
	return server
}

func createRouter(database *ImageDatabase) *http.ServeMux {
	imgSrv := http.FileServer(http.FS(database.root))
	staticSrv := http.FileServer(http.FS(staticFS))
	router := http.NewServeMux()
	router.Handle("/image/", http.StripPrefix("/image/", imgSrv))
	router.Handle("/static/", staticSrv)
	if database == nil {
		return router
	}
	router.HandleFunc("/list", createAPI(database))
	render, err := Browser()
	if err != nil {
		return router
	}
	getBrowser := func(w http.ResponseWriter, r *http.Request) {
		images := database.GetImages(10, 0)
		render(w, images)
	}
	router.HandleFunc("/", getBrowser)
	return router
}

func createAPI(database *ImageDatabase) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, offset := getLimitOffset(r.URL.Query())
		images := database.GetImages(limit, offset)
		err := json.NewEncoder(w).Encode(images)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}

func getLimitOffset(query url.Values) (limit, offset int) {
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}
	offset, err = strconv.Atoi(query.Get("offset"))
	if err != nil {
		offset = 0
	}
	return
}
