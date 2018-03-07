package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	myTemplate, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = myTemplate.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
