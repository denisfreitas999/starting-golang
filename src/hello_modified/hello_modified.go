package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	exibirIntroducao()

	for {
		exibirMenu()

		comando := lerComando()

		switch comando {
			case 1:
				iniciarMonitoramento();
			case 2:
				fmt.Println("Exibindo Logs...")
			case 0:
				fmt.Println("Saindo do Programa...")
				os.Exit(0)
			default:
				fmt.Println("Comando inválido! Escolha entre os valores 0, 1 ou 2 para acessar as opções do menu.")
				os.Exit(-1)
		}
	}
}

func exibirIntroducao() {
	nome := "Denisson"
	var versao float32 = 1.1

	fmt.Println("Olá Sr.", nome)
	fmt.Println("Este programa está na versão:", versao)
}


func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// Define os sites e suas probabilidades
	sites := map[string]float64{
		"https://httpbin.org/status/200": 0.7,
		"https://httpbin.org/status/400": 0.2,
		"https://httpbin.org/status/500": 0.1,
	}

	site := selecionarSite(sites)
	res, err := http.Get(site)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Houve um problema ao carregar o site:", site+". Status code:", res.StatusCode)
	}
}

func selecionarSite(sites map[string]float64) string {
	// Gera um número aleatório entre 0.0 e 1.0
	r := rand.Float64()

	accumulatedProbability := 0.0
	for site, probability := range sites {
		accumulatedProbability += probability
		if r < accumulatedProbability {
			return site
		}
	}

	// Caso algo dê errado, retorna o primeiro site como fallback
	for site := range sites {
		return site
	}

	return ""
}