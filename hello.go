package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramento = 3
const delay = 5

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := lerComando()

		switch comando {
		case 1:
			monitorando()
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
}

func monitorando() {
	for i := 0; i < monitoramento; i++ {
		fmt.Println("Monitorando...")
		sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}
		for i, site := range sites {
			site = sites[i]
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func testaSite(site string) {

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
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
