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

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		Name: "Tony Stark",
		Age:  40,
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
