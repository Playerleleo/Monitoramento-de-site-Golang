package main

import "fmt"

func main() {
	nome := "douglas"
	versao := 1.1
	var resposta int

	fmt.Println("Olá senhor(a)", nome)
	fmt.Println("Este programa esta na versao", versao)

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")

	fmt.Scan(&resposta)

	fmt.Println("A seleção escolhida foi: ", resposta)
}
