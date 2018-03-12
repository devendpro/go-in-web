package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("templateGlobalFunctions2.gohtml"))
}

func main() {

	xs := []string{"zero", "um", "dois", "tres", "quatro", "cinco"}

	data := struct {
		Words []string
		Lnome string
	}{
		xs,
		"Martins",
	}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	err = meuTemplate.Execute(novoArquivo, data)
	if err != nil {
		log.Fatalln(err)
	}

}
