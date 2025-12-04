package config

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
)

type Config struct {
	URL         string
	Requests    int
	Concurrency int
}

func ParseFlags() (*Config, error) {
	config := &Config{}

	flag.StringVar(&config.URL, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&config.Requests, "requests", 0, "Número total de requests")
	flag.IntVar(&config.Concurrency, "concurrency", 0, "Número de chamadas simultâneas")

	flag.Parse()

	return config, config.Validate()
}

func (c *Config) Validate() error {
	if c.URL == "" {
		return errors.New("URL é obrigatória")
	}

	parsedURL, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return fmt.Errorf("URL inválida: %w", err)
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("URL deve usar http ou https")
	}

	if c.Requests <= 0 {
		return errors.New("número de requests deve ser maior que zero")
	}

	if c.Concurrency <= 0 {
		return errors.New("concurrency deve ser maior que zero")
	}

	if c.Concurrency > c.Requests {
		return errors.New("concurrency não pode ser maior que total de requests")
	}

	return nil
}
