// +build !appengine

package main

import (
	"net/http"

	"github.com/captaincodeman/clean-go/adapters/web"
	"github.com/captaincodeman/clean-go/engine"
	"github.com/captaincodeman/clean-go/providers/mongodb"
)

// when running in traditional or 'standalone' mode
// we're going to use MongoDB as the storage provider
// and start the webserver running ourselves.
func main() {
	s := mongodb.NewStorage(config.MongoURL)
	e := engine.NewEngine(s)
	http.ListenAndServe(":8080", web.NewWebAdapter(e, true))
}
