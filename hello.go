package main

import (
	"fmt"
	"os"
)
import _ "os"

func main() {
	exibeIntroducao()
	exibeMenu()
	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("monitorando")
	case 2:
		fmt.Println("exibindo logs")
	case 3:
		fmt.Println("saindo....")
		os.Exit(0)
	default:
		fmt.Println("Algo deu errado")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "douglas"
	versao := 1.1
	fmt.Println("Olá senhor(a)", nome)
	fmt.Println("Este programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("3- Sair do programa")
}

func lerComando() int {
	var resposta int
	fmt.Scan(&resposta)
	fmt.Println("A seleção escolhida foi: ", resposta)
	return resposta
}
func output() {

}
