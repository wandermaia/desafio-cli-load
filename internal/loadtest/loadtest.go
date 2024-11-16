package loadtest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Objeto para armazenamento dos dados dos testes
type Report struct {
	TotalRequests      int
	Status200          int
	StatusDistribution map[int]int
}

func RunLoadTest(url string, totalRequests, concurrency int) Report {

	// Wait group para garantir a execução de todas as goroutines
	var wg sync.WaitGroup

	// Mutex para impedir o race condiction
	var mu sync.Mutex

	// Criando o objeto de report e instanciando o map da distribuição dos status code
	report := Report{StatusDistribution: make(map[int]int)}

	// Calcula a quantidade de requests para cada go rotine
	requestsPerGoroutine := criaArray(totalRequests, concurrency)

	// Gera uma go rotine até atingir o tamanho do array
	for i := range requestsPerGoroutine {
		wg.Add(1)
		go func() {

			defer wg.Done()

			// Executa a chamada na URL a quantidade de vezes de cada gorotine
			for j := 0; j < requestsPerGoroutine[i]; j++ {

				// Definindo o contexto para realizar o timeout em 2 segundos
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
				defer cancel()

				// Cria uma nova requisição HTTP com o contexto
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
				if err != nil {
					log.Fatal(err)
				}

				// Executando a requisição
				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					log.Printf("Erro ao realizar a chamada ao site informado: %s", err)
					continue
				}

				// Bloqueando o objeto de report para inserir os dados.
				mu.Lock()

				//Incrementando a quantidade total de requests
				report.TotalRequests++

				//Incrementando o map de acordo com o status code recebido.
				report.StatusDistribution[resp.StatusCode]++

				// Incrementa o contador exclusivo para código 200
				if resp.StatusCode == 200 {
					report.Status200++
				}
				mu.Unlock()
				fmt.Printf(".")
				resp.Body.Close()
			}
		}()
	}
	wg.Wait()

	// Retornando os dados da requisição.
	return report
}

// Função privada para criar um array que será utilizado da seguinte forma:
// Tamanho do array: Quantidade de goroutines que serão criadas, de acordo com o valor de concorrência solicitado
// Valor em cada posição do array: Número de chamadas que cada go routine vai realizar na URL informada.
func criaArray(dividendo, divisor int) []int {

	// Realizando a divisão inteira e pegando o resto do valor
	quociente := dividendo / divisor
	resto := dividendo % divisor

	// Cria um array com do tamanho do divisor, para representar o paralelismo da requisição
	arr := make([]int, divisor)

	// Inicializa todas as posições com o valor 16
	for i := range arr {
		arr[i] = quociente
	}

	//Adicionando o resto da divisão
	for i := 0; i < resto; i++ {
		arr[i]++
	}

	// Retonando o array com os valores
	return arr

}
