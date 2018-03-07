package main

import (
	"log"
	"os"
	"text/template"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.ParseFiles("template_map.gohtml"))
}

func main() {
	sabios := map[string]string{
		"India":     "Ghanddi",
		"America":   "MLK",
		"Israel":    "Jesus",
		"Palestina": "Maomé"}
	err := meuTemplate.Execute(os.Stdout, sabios)
	if err != nil {
		log.Fatalln(err)
	}
}
