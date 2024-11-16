package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/wandermaia/desafio-cli-load/internal/loadtest"
)

func main() {

	// Coletando os dados de entrada
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")
	flag.Parse()

	// Validando os valores informados.
	if *url == "" || *requests < 1 || *concurrency < 1 {
		log.Fatal("Todos os parâmetros --url, --requests (>= 1) e --concurrency (>= 1) são obrigatórios")
	}

	// Coletando o horário de início da execução
	start := time.Now()

	fmt.Printf("\nTeste em andamento, por favor aguarde a execução das requisições! \n\n")
	report := loadtest.RunLoadTest(*url, *requests, *concurrency)

	//Cálculo do tempo de execução
	duration := time.Since(start)

	fmt.Printf("\n\nTeste finalizado!\n\n")

	// Exbindo os dados das requisições.
	fmt.Printf("Tempo total gasto: %v\n", duration)
	fmt.Printf("Total de requests realizadas: %d\n", report.TotalRequests)
	fmt.Printf("Requests com status 200: %d\n", report.Status200)
	fmt.Printf("Distribuição de status HTTP: %v\n\n", report.StatusDistribution)
}
