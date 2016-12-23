package hello

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	goji.Get("/", handler)
	goji.Get("/user", handlerUser)
	goji.Get("/user/:name", handlerUser)
}

func handler(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "/")
}

func handlerUser(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s/%v", "/user", c.URLParams["name"])
}
