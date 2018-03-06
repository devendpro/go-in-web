package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	nome := "William"

	template := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="pt-br">
	<head>
	<meta charset="UTF-8">
	<title>Olá Mundo</title>
	</head>
	<body>
	<h1>` + nome + `</h1>
	</body>
	</html>
	`)

	arq, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Não foi possível criar o arquivo index.html", err)
	}
	defer arq.Close()

	io.Copy(arq, strings.NewReader(template))
}
