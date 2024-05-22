# multithreading

## Rodando o projeto

Este projeto consiste em um programa em Go que busca o resultado mais rápido entre duas APIs de busca CEP distintas e retorna qual a API mais rapida, limitando o tempo de resposta em 1 segundo. Caso o tempo de resposta ultrapasse esse limite, o programa exibirá um erro de timeout.

### Pré-requisitos

Antes de rodar o projeto, certifique-se de ter o Go instalado em seu sistema. Caso ainda não tenha, você pode baixá-lo e instalá-lo a partir do site oficial do Go: [Aqui](https://golang.org/)

Caso esteja usando windows pode executar o arquivo main.exe


### Execute o programa

Para executar o programa main.go, você deve fornecer o CEP que desaja para consulta como argumentos na linha de comando. Para isso, utilize o seguinte comando:

```
go run main.go -cep XXXXX-XXX
```

Substitua **XXXXX-XXX** pelo cep real que gostaria de testar.

Por exemplo, se você quiser testar o CEP 49400-000 o comando ficaria assim:

```
go run main.go -cep 49400-000
```