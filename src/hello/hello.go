package main

import "fmt"

func main() {
	nome := "Rodrigo"
	versao := 1.1
	fmt.Println("Olá, sr(a). ", nome)
	fmt.Println("Este programa está na versão: ", versao)

	fmt.Println("1- Inicar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O valor da variável comando é: ", comando)
}
