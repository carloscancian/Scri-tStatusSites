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

//Constante de monitoramento seria as vezes de teste que quero realizar
const monitoramento = 3

// Constante delay , seria de quanto em quanto tempo quero testar os sites
const delay = 5

func main() {

	ExibindoIntroducao()
	for {

		ExibeMenu()

		comando := leComando()

		//if comando == 1 {
		//	fmt.Println("Monitorado...")
		//} else if comando == 2 {
		//	fmt.Println("Exibindo logs...")
		//} else if comando == 0 {s
		//	fmt.Println("Terminando...")
		//} else {
		//	fmt.Println("Não conheço este comando")
		//}

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLog()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}

}

func ExibindoIntroducao() {
	nome := "Carlos"
	versao := 1.1
	fmt.Println("Olá, Sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func ExibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir log")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorado...")
	//sites := []string{"https://random-status-code.herokuapp.com/",
	//"https://www.alura.com.br", "https://www.caelum.com.br"}

	//fmt.Println("O meu slace tem tamanho de :", len(sites))

	sites := leSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando Site: ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}

	fmt.Println("")

	//fmt.Println("Digite o Site que deseja monitor")
	//var siteEscolhido string
	//fmt.Scan(&siteEscolhido)
	//fmt.Println(resp)
}

func exibirLog() {
	fmt.Println("Exibindo logs...")
	imprimeLogs()
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:",
			resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "- " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
	}

	fmt.Println(string(arquivo))

}
