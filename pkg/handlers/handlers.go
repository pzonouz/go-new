package handlers

import (
	"fmt"
	"go-new/pkg/renders"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "Not found")
		return

	}
	renders.RenderTemplate(w, "home.tmpl")
}
func About(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "about.tmpl")
}
