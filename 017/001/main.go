package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type police struct {
	Patente string
	Cidade  string
}

func (p police) RetornaCidade() string {
	return p.Cidade + p.Patente
}

func (p person) Someprocessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var meuTemplage *template.Template

func init() {
	meuTemplage = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	p := person{
		Name: "William Martins",
		Age:  40,
	}

	po := police{
		Patente: "Capitão",
		Cidade:  "Montes Claros",
	}

	if po.Cidade == "Minas" {
		log.Fatalln("Minas não é cidade")
	}

	fmt.Println(po.RetornaCidade())

	arq, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer arq.Close()

	err = meuTemplage.Execute(arq, p)
	if err != nil {
		log.Fatalln(err)
	}
}
