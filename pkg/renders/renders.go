package renders

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./web/templates/"+tmpl, "./web/templates/base.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err2 := parsedTemplate.Execute(w, nil)
	if err2 != nil {
		log.Fatal(err2)
	}
}
