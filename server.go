package goart

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type ImageServer struct {
	http.Handler
}

func CreateImageServer(database *ImageDatabase) *ImageServer {
	server := &ImageServer{}
	router := createRouter(database)
	server.Handler = router
	return server
}

func createRouter(database *ImageDatabase) *http.ServeMux {
	imgSrv := http.FileServer(http.FS(database.root))
	router := http.NewServeMux()
	router.Handle("/static/", http.StripPrefix("/static/", imgSrv))
	getList := func(w http.ResponseWriter, r *http.Request) {
		limit, offset := getLimitOffset(r.URL.Query())
		json.NewEncoder(w).Encode(database.GetImages(limit, offset))
	}
	router.HandleFunc("/list", getList)
	return router
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
