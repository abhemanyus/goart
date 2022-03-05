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

func Browser() (func(w io.Writer, images ImageList) error, error) {
	templ, err := template.ParseFS(browserTemplate, "templates/browser.html")
	if err != nil {
		return nil, err
	}

	return func(w io.Writer, images ImageList) error {
		err = templ.ExecuteTemplate(w, "browser.html", images)
		return nil
	}, nil
}
