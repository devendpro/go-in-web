package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("templateGlobalFunctions.gohtml"))
}

func main() {

	xs := []string{"zero", "um", "dois", "tres", "quatro", "cinco"}

	err := meuTemplate.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
