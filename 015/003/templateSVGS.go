package main

import (
	"html/template"
	"log"
	"net/http"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)

	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	err := meuTemplate.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		log.Fatalln("Não foi possível criar os templates HTML.", err)
	}

}
