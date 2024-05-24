package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type Address struct {
	CEP        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	UF         string `json:"uf"`
}

type BrasilAPIAddress struct {
	CEP          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

type APIResult struct {
	Address Address
	Source  string
}

func fetchFromBrasilAPI(cep string, ch chan<- APIResult, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var brasilAPIAddress BrasilAPIAddress
	if err := json.Unmarshal(body, &brasilAPIAddress); err != nil {
		return
	}

	address := Address{
		CEP:        brasilAPIAddress.CEP,
		Logradouro: brasilAPIAddress.Street,
		Bairro:     brasilAPIAddress.Neighborhood,
		Localidade: brasilAPIAddress.City,
		UF:         brasilAPIAddress.State,
	}

	ch <- APIResult{Address: address, Source: "BrasilAPI"}
}

func fetchFromViaCEP(cep string, ch chan<- APIResult, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var address Address
	if err := json.Unmarshal(body, &address); err != nil {
		return
	}

	ch <- APIResult{Address: address, Source: "ViaCEP"}
}

func main() {
	cep := flag.String("cep", "", "O CEP a ser buscado (formato XXXXX-XXX)")
	flag.Parse()

	if *cep == "" {
		fmt.Println("Erro: Você deve fornecer um CEP usando o parâmetro -cep")
		os.Exit(1)
	}

	ch := make(chan APIResult, 2)
	var wg sync.WaitGroup

	wg.Add(2)
	go fetchFromBrasilAPI(*cep, ch, &wg)
	go fetchFromViaCEP(*cep, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case result := <-ch:
		fmt.Printf("Resultado da API %s:\n", result.Source)
		fmt.Printf("CEP: %s\nLogradouro: %s\nBairro: %s\nLocalidade: %s\nUF: %s\n", result.Address.CEP, result.Address.Logradouro, result.Address.Bairro, result.Address.Localidade, result.Address.UF)
	case <-time.After(1 * time.Second):
		fmt.Println("Erro: Timeout")
	}
}
