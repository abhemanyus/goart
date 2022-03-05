package goart

import (
	"embed"
	"net/http"
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
	render, err := Browser()
	if err != nil {
		return router
	}
	getBrowser := func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil {
			page = 0
		}
		images := database.GetImages(10, 10*page)
		if len(images) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if len(images) < 10 {
			page = -1
		}
		render(w, images, page)
	}
	router.HandleFunc("/", getBrowser)
	return router
}
