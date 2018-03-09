package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var meuTemplate *template.Template

type sabio struct {
	Nome string
	Lema string
}

// Criar uma FuncMap para registrar as funções
// "uc" é a função que irá fazer chamada dentro do template
// "uc" é a função que muda uma string para caixa alta.
// "ft" é a função declarada por mim
// "ft" retira uma parte de uma string, retornando os 3 primeiros caracteres.

var funcaoMap = template.FuncMap{
	"uc":   strings.ToUpper,
	"str3": primeiros3Caracteres,
}

func init() {
	meuTemplate = template.Must(template.New("").Funcs(funcaoMap).ParseFiles("templateFunc.gohtml"))
}

func primeiros3Caracteres(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func main() {
	Sabios := []sabio{
		{
			Nome: "Jesus",
			Lema: "Amar a Deus acima de todas as coisas",
		},
		{
			Nome: "Martin Luther King",
			Lema: "Todos tem direitos iguais",
		},
	}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer novoArquivo.Close()

	err = meuTemplate.ExecuteTemplate(novoArquivo, "templateFunc.gohtml", Sabios)
	if err != nil {
		log.Fatalln(err)
	}
}
