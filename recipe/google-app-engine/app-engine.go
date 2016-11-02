// +build appengine

package main

import (
	"github.com/insionng/vodka"
	"github.com/insionng/vodka/engine/standard"
	"net/http"
)

func createMux() *vodka.Vodka {
	e := vodka.New()

	// note: we don't need to provide the middleware or static handlers, that's taken care of by the platform
	// app engine has it's own "main" wrapper - we just need to hook vodka into the default handler
	s := standard.New("")
	s.SetHandler(e)
	http.Handle("/", s)

	return e
}
