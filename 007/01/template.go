package main

import (
	"fmt"
)

func main() {
	nome := "William"

	template := `
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
	`

	fmt.Println(template)
}
