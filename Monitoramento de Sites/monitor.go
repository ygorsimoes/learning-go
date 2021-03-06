package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	fmt.Println("|===============================|")
	fmt.Println("|     MONITORAMENTO DE SITES    |")
	fmt.Println("|===============================|")

	fmt.Println("\n1 - Iniciar Monitoramento;")
	fmt.Println("2 - Exibir Logs;")
	fmt.Println("0 - Sair do programa;")

	var comandoLido int
	fmt.Print("\nOpção >: ")
	fmt.Scan(&comandoLido)

	comando := comandoLido

	switch comando {
	case 1:
		fmt.Print("\n[+] Iniciando Monitoramento...\n\n")
		iniciarMonitoramento()
	case 2:
		fmt.Println("\n[+] Exibindo Logs...")
		imprimeLogs()
	case 3:
		fmt.Println("\n[:)] Saindo... Até mais!")
		os.Exit(0)
	default:
		fmt.Println("\n[!] Comando não reconhecido.")
		fmt.Println("[!] Caso o erro persista contate o desenvolvedor.")
		os.Exit(-1)
	}

}

func iniciarMonitoramento() {

	sites := leSitesDoArquivo()

	for true {

		// Percorre a lista de sites
		for _, site := range sites {

			// Realiza a requisição web
			resp, err := http.Get(site)

			if err != nil {
				fmt.Println("[!] Ocorreu um erro inesperado!", err)
			}

			// Verifica o Status Code da requisição realizada
			if resp.StatusCode == 200 {
				fmt.Println(resp.StatusCode, "[ON]", site)
				registraLog(site, true)
			} else {
				fmt.Println(resp.StatusCode, "[OF]", site)
				registraLog(site, false)
			}
		}
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	// Abre o arquivo sites.txt
	arquivo, err := os.Open("sites.txt")


	// Trata o possível erro do arquivo
	if err != nil {
		fmt.Println("[!] Ocorreu um erro inesperado!", err)
		os.Exit(-1)
	}

	// Possibilita manipular o arquivo com métodos de leitura
	leitor := bufio.NewReader(arquivo)

	for true {
		// Lê uma linha do arquivo
		linha, err := leitor.ReadString('\n')
		// Remove espaços e quebras de linhas
		linha = strings.TrimSpace(linha)

		// Condição caso não tenha mais conteúdo para leitura no arquivo
		if err == io.EOF {
			break
		}

		// Adiciona a linha dentro do slices
		sites = append(sites, linha)
	}

	// Fecha o arquivo sites.txt
	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {

	// Cria um arquivo chamado "log.txt".
	// Flag: os.O_RDWR   -> Leitura e escrita do arquivo;
	// Flag: os.O_CREATE -> Cria o arquivo caso ele não exista;
	// Flag: os.O_APPEND -> Escreve na última linha do arquivo;
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)


	// Trata o possível erro do arquivo
	if err != nil {
		fmt.Println("[!] Ocorreu um erro inesperado!", err)
		os.Exit(-1)
	}

	// Pega o horário atual formatado
	horarioAtual := time.Now().Format("02/01/2006 15:04:05")

	// Escreve no arquivo logs.txt
	arquivo.WriteString(horarioAtual + " - " + site + " - ONLINE: " + strconv.FormatBool(status) + "\n")


	// Fecha o arquivo log.txt
	arquivo.Close()
}

func imprimeLogs() {
	// Lê o arquivo de logs e depois fecha
	// (Modelo mais rápido de lê um arquivo e depois fecha-lo)
	arquivo, err := ioutil.ReadFile("log.txt")

	// Trata o possível erro do arquivo
	if err != nil {
		fmt.Println("[!] Ocorreu um erro inesperado!", err)
		os.Exit(-1)
	}

	// Coverte o arquivo de bytes para uma string legível
	// Imprime o arquivo log.txt na tela
	fmt.Println(string(arquivo))
}