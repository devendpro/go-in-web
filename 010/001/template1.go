package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("template_variavel.gohtml"))
}

func main() {

	err := meuTemplate.ExecuteTemplate(os.Stdout, "template_variavel.gohtml", `Não foque em si mesmo; Foque na vida como um todo.`)
	if err != nil {
		log.Fatalln(err)
	}

}
