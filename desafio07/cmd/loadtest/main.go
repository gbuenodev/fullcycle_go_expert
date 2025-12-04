package main

import (
	"fmt"
	"os"

	"github.com/gbuenodev/fullcycle_go_expert/desafio07/internal/config"
	"github.com/gbuenodev/fullcycle_go_expert/desafio07/internal/reporter"
	"github.com/gbuenodev/fullcycle_go_expert/desafio07/internal/tester"
)

func main() {
	cfg, err := config.ParseFlags()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Executar teste
	fmt.Printf("Iniciando teste de carga...\n")
	fmt.Printf("URL: %s\n", cfg.URL)
	fmt.Printf("Total de requests: %d\n", cfg.Requests)
	fmt.Printf("Concorrência: %d\n", cfg.Concurrency)
	fmt.Println()

	stats := tester.RunLoadTest(cfg.URL, cfg.Requests, cfg.Concurrency)

	// Exibir relatório
	reporter.PrintReport(stats)
}
