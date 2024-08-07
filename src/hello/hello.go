package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 3
const delay = 2

func main() {
	exibirIntroducao()

	for {
		exibirMenu()
		
		comando := lerComando()

		switch comando {
			case 1:
				iniciarMonitoramento();
			case 2:
				imprimeLogs()
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
	
	sites := lerSitesDoArquivo()
	
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, "Site:", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
	
	fmt.Println("")
}


func testaSite(site string) {

    resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro detectado: ", err)
	}

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
		registrarLog(site, true)
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registrarLog(site, false)
    }
}

	func lerSitesDoArquivo() []string {
		var sites []string

		arquivo, err := os.Open("sites.txt")

		if err != nil {
			fmt.Println("Erro detectado: ", err)
			return sites
		}

		leitor := bufio.NewReader(arquivo)
		
		for {
			linha, err := leitor.ReadString('\n')
			linha = strings.TrimSpace(linha)
			
			sites = append(sites, linha)

			if err == io.EOF {
				break
			}
		}

		arquivo.Close()

		return sites
	}

	func registrarLog(site string, status bool) {
        arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

        if err != nil {    
            fmt.Println("Erro detectado:", err)
            return  // Encerra a função em caso de erro
        }

        // Escreve a data e hora atual, o site e o status online no arquivo
        _, err = arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + 
                                     " | Site: " + site + 
                                     " | Status Online: " + strconv.FormatBool(status) + "\n")

        if err != nil {
            fmt.Println("Erro ao escrever no arquivo:", err)
        }

        arquivo.Close()
    }

	func imprimeLogs() {
		fmt.Println("Exibindo Logs...")
		arquivo, err := os.ReadFile("logs.txt")

		if err != nil {
			fmt.Println("Erro detectado:", err)
		}

		fmt.Println(string(arquivo))
	}