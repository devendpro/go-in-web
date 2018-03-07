package main

import (
	"html/template"
	"log"
	"os"
)

var meuTemplate *template.Template

type sabio struct {
	Nome string
	Lema string
}

func init() {
	meuTemplate = template.Must(template.ParseFiles("template_struct2.gohtml"))
}

func main() {
	buddha := sabio{
		Nome: "Buddha",
		Lema: "Paz",
	}

	gandhi := sabio{
		Nome: "Ghandi",
		Lema: "Esperança",
	}

	mlk := sabio{
		Nome: "Martin Luther King",
		Lema: "Direitos civis",
	}

	jesus := sabio{
		Nome: "Jesus",
		Lema: "Amor",
	}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	sabios := []sabio{buddha, gandhi, mlk, jesus}
	err = meuTemplate.Execute(novoArquivo, sabios)
	if err != nil {
		log.Fatalln(err)
	}

}
