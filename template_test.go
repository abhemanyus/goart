package goart_test

import (
	"bytes"
	"testing"

	"github.com/abhemanyus/goart"
	approvals "github.com/approvals/go-approval-tests"
)

func TestTemplating(t *testing.T) {
	t.Run("convert ImageList into html", func(t *testing.T) {
		imageList := goart.ImageList{
			"picOne.jpg",
			"picTwo.png",
			"picThree.webp",
		}
		buf := bytes.Buffer{}
		render, err := goart.Browser()
		assertError(t, err, nil)
		err = render(&buf, imageList)
		assertError(t, err, nil)
		approvals.VerifyString(t, buf.String())
	})
}
