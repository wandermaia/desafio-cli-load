# Sistema de Stress test

Este repositório foi criado para hospedar o código de implementação do sistema CLI em Go para realizar testes de carga em um serviço web.


## Descrição do Desafio

A seguir estão os dados fornecidos na descrição do desafio.


### Objetivo

Criar um sistema CLI em Go para realizar testes de carga em um serviço web. O usuário deverá fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas.

O sistema deverá gerar um relatório com informações específicas após a execução dos testes.


### Entrada de Parâmetros via CLI

**--url:** URL do serviço a ser testado.

**--requests:** Número total de requests.

**--concurrency:** Número de chamadas simultâneas.

### Execução do Teste


 - Realizar requests HTTP para a URL especificada.

 - Distribuir os requests de acordo com o nível de concorrência definido.

 - Garantir que o número total de requests seja cumprido.


### Geração de Relatório

- Apresentar um relatório ao final dos testes contendo:
    - Tempo total gasto na execução
    - Quantidade total de requests realizados.
    - Quantidade de requests com status HTTP 200.
    - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).


### Execução da aplicação

- Poderemos utilizar essa aplicação fazendo uma chamada via docker. Ex:

```bash

docker run <sua_imagem_docker> —url=http://google.com —requests=1000 —concurrency=10

```

# Execução do Desafio

O sistema utiliza a seguinte hierarquia de arquivos e pastas:

```bash

wander@bsnote283:~desafio-cli-load$ tree
.
├── cmd
│   └── main.go
├── Dockerfile
├── go.mod
├── internal
│   └── loadtest
│       ├── loadtest.go
│       └── loadtest_test.go
├── LICENSE
└── README.md

3 directories, 7 files
wander@bsnote283:~desafio-cli-load$ 


```

A seguir estão descritos como executar os testes automatizados e a aplicação (através de container docker).

## Testes automatizados

Para a realização dos testes, basta executar o seguinte comando `go test ./...` a partir da raiz dos módulos. Abaixo segue o exemplo da execução:


```bash

wander@bsnote283:~desafio-cli-load$ go test ./...
?   	github.com/wandermaia/desafio-cli-load/cmd	[no test files]
ok  	github.com/wandermaia/desafio-cli-load/internal/loadtest	1.329s
wander@bsnote283:~desafio-cli-load$ 


```

## Execução da Aplicação

Para executar o projeto, devemos primeiramente criar a imagem de container, utilizando o comando docker seguir a partir da pasta raiz do projeto:

```bash

wander@bsnote283:~desafio-cli-load$ docker build -t cli-load .
[+] Building 10.3s (10/10) FINISHED                                                                                                  docker:default
 => [internal] load build definition from Dockerfile                                                                                           0.0s
 => => transferring dockerfile: 174B                                                                                                           0.0s
 => [internal] load metadata for docker.io/library/golang:1.23.3-alpine3.20                                                                    1.1s
 => [internal] load .dockerignore                                                                                                              0.0s
 => => transferring context: 2B                                                                                                                0.0s
 => [1/5] FROM docker.io/library/golang:1.23.3-alpine3.20@sha256:c694a4d291a13a9f9d94933395673494fc2cc9d4777b85df3a7e70b3492d3574              0.0s
 => [internal] load build context                                                                                                              0.0s
 => => transferring context: 7.18kB                                                                                                            0.0s
 => CACHED [2/5] WORKDIR /app                                                                                                                  0.0s
 => [3/5] COPY . .                                                                                                                             0.1s
 => [4/5] RUN go mod tidy                                                                                                                      0.5s
 => [5/5] RUN go build -o load-tester ./cmd                                                                                                    8.3s
 => exporting to image                                                                                                                         0.2s
 => => exporting layers                                                                                                                        0.2s
 => => writing image sha256:f4936e1f3e07ab50eb3bcf89ef393a2f9d19d226b1021c45ec9ff68fd7166c81                                                   0.0s
 => => naming to docker.io/library/cli-load                                                                                                    0.0s
wander@bsnote283:~desafio-cli-load$ 

```

Após a execução do comando anterior, já teremos a imagem e podemos executar a aplicação através do docker, confomre o comando abaixo:


```bash

wander@bsnote283:~desafio-cli-load$ docker run cli-load --url=https://google.com --requests=10 --concurrency=4

Teste em andamento, por favor aguarde a execução das requisições! 

..........

Teste finalizado!

Tempo total gasto: 850.472357ms
Total de requests realizadas: 10
Requests com status 200: 10
Distribuição de status HTTP: 

Status Code: 200, Quantidade de ocorrências: 10

wander@bsnote283:~desafio-cli-load$ 


```

Para este exemplo, foi utilizado o comando `docker run cli-load --url=https://google.com --requests=10 --concurrency=4`



