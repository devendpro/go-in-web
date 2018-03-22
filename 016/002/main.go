package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenceToKill bool
}

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p1 := doubleZero{
		person{
			Name: "Capitão Nascimento",
			Age:  45,
		},
		true,
	}

	arq, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer arq.Close()

	err = meuTemplate.Execute(arq, p1)
	if err != nil {
		log.Fatalln(err)
	}
}
