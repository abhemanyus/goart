package goart

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/browser.html"
	browserTemplate embed.FS
)

func Browser() (func(w io.Writer, images ImageList, page int) error, error) {
	templ, err := template.ParseFS(browserTemplate, "templates/browser.html")
	if err != nil {
		return nil, err
	}

	return func(w io.Writer, images ImageList, page int) error {
		err = templ.ExecuteTemplate(w, "browser.html", struct {
			List ImageList
			Next int
			End  bool
		}{images, page + 1, page == -1})
		if err != nil {
			return err
		}
		return nil
	}, nil
}
