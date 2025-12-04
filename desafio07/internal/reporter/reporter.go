package reporter

import (
	"fmt"
	"strings"

	"github.com/gbuenodev/fullcycle_go_expert/desafio07/internal/tester"
)

func PrintReport(stats *tester.Statistics) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("                 LOAD TEST REPORT")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	fmt.Printf("Tempo total de execução: %v\n", stats.TotalDuration)
	fmt.Println()

	fmt.Printf("Total de requests realizados: %d\n", stats.TotalRequests)
	fmt.Println()

	fmt.Printf("Requests com status 200: %d (%.2f%%)\n",
		stats.SuccessCount,
		percentage(stats.SuccessCount, stats.TotalRequests),
	)
	fmt.Println()

	fmt.Println("Distribuição de outros códigos de status:")
	if len(stats.StatusCodeDist) == 0 {
		fmt.Println("  (nenhum)")
	} else {
		for code, count := range stats.StatusCodeDist {
			if code == 200 {
				continue // Já mostramos os 200 acima
			}
			fmt.Printf("  Status %d: %d requests (%.2f%%)\n",
				code,
				count,
				percentage(count, stats.TotalRequests),
			)
		}
	}
	fmt.Println()

	if stats.ErrorCount > 0 {
		fmt.Printf("Erros de rede/timeout: %d (%.2f%%)\n",
			stats.ErrorCount,
			percentage(stats.ErrorCount, stats.TotalRequests),
		)
		fmt.Println()
	}

	fmt.Println(strings.Repeat("=", 60))
}

func percentage(part, total int) float64 {
	if total == 0 {
		return 0
	}
	return float64(part) / float64(total) * 100
}
