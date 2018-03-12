package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("templateGlobalFunctions3.gohtml"))
}

type usuario struct {
	Nome  string
	Lema  string
	Admin bool
}

func main() {
	u1 := usuario{
		Nome:  "Jesus",
		Lema:  "O Amor de Deus nos faz um",
		Admin: true,
	}
	u2 := usuario{
		Nome:  "Martin Luther King",
		Lema:  "Direitos civis",
		Admin: false,
	}
	u3 := usuario{
		Nome:  "",
		Lema:  "Ninguém",
		Admin: true,
	}

	usuarios := []usuario{u1, u2, u3}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	err = meuTemplate.Execute(novoArquivo, usuarios)
	if err != nil {
		log.Fatalln(err)
	}
}
