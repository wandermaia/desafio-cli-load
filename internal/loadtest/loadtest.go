package loadtest

import (
	"fmt"
	"net/http"
	"sync"
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

	//fmt.Println("Preparando as informações para iniciar o teste.")
	// Calcula a quantidade de requests para cada go rotine
	requestsPerGoroutine := criaArray(totalRequests, concurrency)

	//fmt.Println("Iniciando o teste. Dependendo da quantidade de chamadas solicitadas, será necessário aguardar um pouco")

	// Gera uma go rotine até atingir o tamanho do array
	for i := range requestsPerGoroutine {
		wg.Add(1)
		go func() {

			defer wg.Done()

			//log.Printf("Go rotine %v", i)

			// Executa a chamada na URL a quantidade de vezes de cada gorotine
			for j := 0; j < requestsPerGoroutine[i]; j++ {
				//for j := 0; j < requestsPerGoroutine; j++ {

				//	log.Printf("Go rotine %v, Execução %v", i, j)
				resp, err := http.Get(url)
				if err != nil {
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

	// fmt.Printf("\n\nFunção do Array\n\n")
	// fmt.Printf("Dividendo %d, Divisor: %d\n", dividendo, divisor)
	// fmt.Printf("Quociente: %d, Resto: %d\n\n", quociente, resto)

	// Cria um array com do tamanho do divisor, para representar o paralelismo da requisição
	arr := make([]int, divisor)

	// Inicializa todas as posições com o valor 16
	for i := range arr {
		arr[i] = quociente
	}

	// Imprime o array sem o resto da divisão
	// fmt.Println(arr)

	//Adicionando o resto da divisão
	for i := 0; i < resto; i++ {
		arr[i]++
	}

	// fmt.Println("Array alterado:")
	// fmt.Println(arr)

	// Retonando o array com os valores
	return arr

}
