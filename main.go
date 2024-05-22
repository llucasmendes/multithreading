package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"time"
)

type MessageReturn struct {
	Message  string `json:"response-time"`
	Response string `json:"response"`
}

type ApiReturn struct {
	Message string `json:"message"`
	Body    []byte `json:"response"`
}

func main() {

	cep := flag.String("cep", "", "URL da primeira API")
	flag.Parse()

	urlAPI1 := "https://cdn.apicep.com/file/apicep/" + *cep + ".json"
	urlAPI2 := "http://viacep.com.br/ws/" + *cep + "/json/"
	timeout := 1 * time.Second

	ch := make(chan MessageReturn)
	ch2 := make(chan MessageReturn)

	go worker(ch, urlAPI1)
	go worker(ch2, urlAPI2)

	select {
	case result := <-ch:
		value, _ := json.Marshal(result)
		fmt.Println(string(value))
	case result := <-ch2:
		value, _ := json.Marshal(result)
		fmt.Println(string(value))
	case <-time.After(timeout):
		fmt.Println("Erro: timeout ao aguardar a resposta.")
		return
	}

}

func worker(ch chan<- MessageReturn, url string) {

	apiReturn, err := fetchAPI(url)
	if err != nil {
		fmt.Printf("Erro: %v \n", err)
		return
	}

	messageReturn := MessageReturn{
		Message:  apiReturn.Message,
		Response: string(apiReturn.Body),
	}

	ch <- messageReturn
}

func fetchAPI(url string) (ApiReturn, error) {
	client := http.Client{}
	startTime := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		return ApiReturn{}, err
	}
	defer resp.Body.Close()
	elapsedTime := time.Since(startTime)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ApiReturn{}, err
	}
	message := fmt.Sprintf("Tempo de resposta para %s: %s", url, elapsedTime)

	return ApiReturn{
		message,
		body,
	}, nil
}
