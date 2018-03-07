package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

type sabio struct {
	Nome string
	Lema string
}

func init() {
	meuTemplate = template.Must(template.ParseFiles("template_struct.gohtml"))
}

func main() {
	jesus := sabio{
		Nome: "Jesus",
		Lema: "Amar a Deus acima de todas as coisas, e ao teu próximo como a ti mesmo.",
	}

	err := meuTemplate.Execute(os.Stdout, jesus)
	if err != nil {
		log.Fatalln(err)
	}
}
