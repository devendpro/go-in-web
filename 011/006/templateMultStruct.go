package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("templateMultStruct.gohtml"))
}

type sabio struct {
	Nome string
	Lema string
}

type carro struct {
	Marca  string
	Modelo string
	Portas int
}

type itens struct {
	Sabios []sabio
	Carros []carro
}

func main() {
	ossabios := []sabio{
		{
			Nome: "Jesus",
			Lema: "Amor",
		},
		{
			Nome: "Martin L. King",
			Lema: "Direitos civis",
		},
	}

	oscarros := []carro{
		{
			Marca:  "Ford",
			Modelo: "New Fiesta",
			Portas: 4,
		},
		{
			Marca:  "GM",
			Modelo: "Camaro",
			Portas: 2,
		},
	}

	data := itens{
		Sabios: ossabios,
		Carros: oscarros,
	}

	meuArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer meuArquivo.Close()

	err = meuTemplate.Execute(meuArquivo, data)
	if err != nil {
		log.Fatalln(err)
	}

}
