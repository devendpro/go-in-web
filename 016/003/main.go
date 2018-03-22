package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fail, Spring, Summer semester
}

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fail: semester{
			Term: "Fail",
			Courses: []course{
				course{"ABC", "Introdução a progrmação", "1"},
				course{"CC1", "Calculo 1", "1"},
				course{"CC2", "Calculo 2", "1"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"XYZ", "Programação avançada", "2"},
				course{"XYZB", "Banco de dados 1", "2"},
				course{"XYZB", "Banco de dados 2", "2"},
			},
		},
	}

	arq, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer arq.Close()

	err = meuTemplate.Execute(arq, y)
	if err != nil {
		log.Fatalln(err)
	}
}
