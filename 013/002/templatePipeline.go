package main

import (
	"html/template"
	"log"
	"math"
	"os"
)

var meuTemplate *template.Template

func init() {
	meuTemplate = template.Must(template.New("").Funcs(funcoes).ParseFiles("templatePipeline.gohtml"))
}

func dobro(x int) int {
	return 2 * x
}

func potencia(x int) float64 {
	return math.Pow(float64(x), 2)
}

func raizQuadrada(x float64) float64 {
	return math.Sqrt(x)
}

var funcoes = template.FuncMap{
	"fdobro":        dobro,
	"fpotencia":     potencia,
	"fraizQuadrada": raizQuadrada,
}

func main() {

	novoArquivo, err := os.Create("index.html")
	if err != nil {
		log.Fatalf("Ocorreu um erro. O erro é: %s", err)
	}
	defer novoArquivo.Close()

	meuTemplate.ExecuteTemplate(novoArquivo, "templatePipeline.gohtml", 3)
}
