package hello

import (
	"html/template"
	"io"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func init() {
	http.Handle("/", goji.DefaultMux)

	goji.Get("/", homeHandler)
	goji.Get("/user", userHandler)
	goji.Get("/user/:name", userHandler)
}

func render(v string, w io.Writer, data map[string]interface{}) {
	tmpl := template.Must(template.ParseFiles("views/layout.html", v))
	tmpl.Execute(w, data)
}

func homeHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Name": "home",
	}
	render("views/home.html", w, data)
}

func userHandler(c web.C, w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Name": c.URLParams["name"],
	}
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "NAME: "+c.URLParams["name"])

	render("views/user.html", w, data)
}
