package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("template_struct3.gohtml"))
}

type sabio struct {
	Nome string
	Lema string
}

func main() {
	jesus := sabio{
		Nome: "Jesus",
		Lema: "Amor",
	}

	mlk := sabio{
		Nome: "Martin Luther King",
		Lema: "Direitos Civis",
	}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	sabios := []sabio{jesus, mlk}

	err = meuTemplate.Execute(novoArquivo, sabios)
	if err != nil {
		log.Fatalln(err)
	}

}
