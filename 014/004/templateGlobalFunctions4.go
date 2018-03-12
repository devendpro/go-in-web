package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("templateGlobalFunctions4.gohtml"))
}

func main() {

	pontuacao := struct {
		Placar1 int
		Placar2 int
	}{
		7,
		9,
	}
	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	err = meuTemplate.Execute(novoArquivo, pontuacao)
	if err != nil {
		log.Fatalln(err)
	}
}
