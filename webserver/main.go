package main

import (
	"log"
	"net/http"
	"os"

	"github.com/abhemanyus/goart"
)

func main() {
	f := os.DirFS(os.Args[1])
	database, err := goart.CreateImageStore(f)
	if err != nil {
		log.Fatal(err)
	}
	server := goart.CreateImageServer(database)
	log.Fatal(http.ListenAndServe(":"+os.Args[2], server))
}
