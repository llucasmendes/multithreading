
# Multithreaded CEP Lookup in Go

Este projeto demonstra como realizar requisições simultâneas a duas APIs de CEP (Código de Endereçamento Postal) em Go, retornando o resultado mais rápido.

## Descrição

O programa faz requisições simultâneas para as seguintes APIs:

1. [BrasilAPI](https://brasilapi.com.br/api/cep/v1/)
2. [ViaCEP](http://viacep.com.br/ws/)

A primeira resposta recebida será exibida no terminal, enquanto a resposta mais lenta será descartada. Se nenhuma resposta for recebida em 1 segundo, uma mensagem de timeout será exibida.

## Requisitos

- Go 1.16 ou superior

## Como Usar

### Passo 1: Clonar o Repositório

Clone o repositório para sua máquina local:

```bash
git clone https://github.com/seu-usuario/multithreaded-cep-lookup.git
cd multithreaded-cep-lookup
```

### Passo 2: Executar o Programa

Para executar o programa main.go, você deve fornecer o CEP que desaja para consulta como argumentos na linha de comando. Para isso, utilize o seguinte comando:

```
go run main.go -cep XXXXX-XXX
```

Por padrão, o CEP usado no exemplo é `01153000`. Você pode alterar o valor do CEP diretamente no código (`main.go`) se desejar testar com outro CEP.

### Exemplo de Saída

```
go run main.go -cep 49035-655
```


```plaintext
Resultado da API ViaCEP:
CEP: 49035-655
Logradouro: Avenida Conselheiro João Moreira Filho
Bairro: Coroa do Meio
Localidade: Aracaju
UF: SE
```

Ou, em caso de timeout:

```plaintext
Erro: Timeout
```

## Explicação do Código

### Estruturas de Dados

- `Address`: Define a estrutura dos dados do endereço.
- `APIResult`: Define a estrutura para armazenar o resultado da API e sua origem.

### Funções

- `fetchFromBrasilAPI(cep string, ch chan<- APIResult, wg *sync.WaitGroup)`: Faz a requisição à BrasilAPI.
- `fetchFromViaCEP(cep string, ch chan<- APIResult, wg *sync.WaitGroup)`: Faz a requisição à ViaCEP.

### Função Principal (`main`)

1. Define o CEP a ser consultado (`cep := "01153000"`).
2. Cria um canal (`ch`) para receber os resultados.
3. Utiliza `sync.WaitGroup` para aguardar ambas as goroutines terminarem.
4. Inicia as goroutines para fazer as requisições às APIs simultaneamente.
5. Usa `select` para esperar a primeira resposta ou um timeout de 1 segundo.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

