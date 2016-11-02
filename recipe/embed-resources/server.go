package main

import (
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
)

func main() {
	e := vodka.New()
	// the file server for rice. "app" is the folder where the files come from.
	assetHandler := http.FileServer(rice.MustFindBox("app").HTTPBox())
	// serves the index.html from rice
	e.GET("/", standard.WrapHandler(assetHandler))

	// servers other static files
	e.GET("/static/*", standard.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	e.Run(standard.New(":3000"))
}
