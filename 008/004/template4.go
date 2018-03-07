package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := meuTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = meuTemplate.ExecuteTemplate(os.Stdout, "um.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = meuTemplate.ExecuteTemplate(os.Stdout, "dois.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = meuTemplate.ExecuteTemplate(os.Stdout, "tres.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
