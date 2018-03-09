package main

import (
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.New("").Funcs(funcaoTemplate).ParseFiles("templateTime.gohtml"))
}

func diaMesAno(t time.Time) string {
	return t.Format("02/01/2006")
}

func caixaAlta(s string) string {
	s = strings.ToUpper(s)
	return s
}

var funcaoTemplate = template.FuncMap{
	"fDateMDY":  diaMesAno,
	"caixaAlta": caixaAlta,
}

func main() {

	type dados struct {
		Datahora time.Time
		Frase    string
	}

	d1 := []dados{
		{
			time.Now(),
			"William de Jesus Martins",
		},
		{
			time.Now(),
			"Paule Christe",
		},
	}

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Ocorreu um erro: Erro %s", err)
	}

	err = meuTemplate.ExecuteTemplate(novoArquivo, "templateTime.gohtml", d1)
	if err != nil {
		log.Fatalf("Ocorreu um erro. Erro: %s", err)
	}
}
