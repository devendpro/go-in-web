package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseGlob("template/*.gohtml"))
}

func main() {
	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	err = meuTemplate.ExecuteTemplate(novoArquivo, "index.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
