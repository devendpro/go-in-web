package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	meuTemplate, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Println("Ocorreu um erro ao tentar criar o arquivo index.html.", err)
	}
	defer novoArquivo.Close()

	err = meuTemplate.Execute(novoArquivo, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
