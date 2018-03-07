package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {

	meuTemplate = template.Must(template.ParseFiles("template_slice_com_indice.gohtml"))
}

func main() {

	sabios := []string{"Jesus", "MLK", "Gandhi", "Mandela"}

	meuArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer meuArquivo.Close()

	err = meuTemplate.Execute(meuArquivo, sabios)
	if err != nil {
		log.Fatalln(err)
	}

}
