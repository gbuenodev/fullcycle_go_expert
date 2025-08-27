package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func consultaCep(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GET: %s, status: %d, response: %s", url, resp.StatusCode, string(body))
	}
	return io.ReadAll(resp.Body)
}

func main() {

	var cep string

	fmt.Print("Digite o CEP: ")
	fmt.Scanln(&cep)

	api1 := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	api2 := fmt.Sprintf("http://brasilapi.com.br/api/cep/v1/%s", cep)

	ctx, timeout := context.WithTimeout(context.Background(), 1*time.Second)
	defer timeout()

	type res struct {
		api  string
		body []byte
		err  error
	}

	ch := make(chan res, 2)

	for _, api := range []string{api1, api2} {
		go func(api string) {
			body, err := consultaCep(ctx, api)
			ch <- res{api, body, err}
		}(api)
	}

	select {
	case res := <-ch:
		if res.err != nil {
			fmt.Printf("Erro na resposta da API %s: %v\n", res.api, res.err)
		} else {
			fmt.Printf("Primeira resposta recebida da API %s: %s\n", res.api, string(res.body))
		}
		// Descarta a resposta da outra goroutine para evitar vazamento
		<-ch
	case <-ctx.Done():
		fmt.Println("Tempo esgotado: nenhuma API respondeu a tempo")
	}
}
