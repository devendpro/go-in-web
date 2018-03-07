package main

import (
	"html/template"
	"log"
	"os"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	err := meuTemplate.ExecuteTemplate(os.Stdout, "template.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
