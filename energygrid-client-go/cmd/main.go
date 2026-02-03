package main

import (
	"fmt"

	"energygrid-client-go/internal/limiter"
	"energygrid-client-go/internal/utils"
)

func main() {
	fmt.Println(" EnergyGrid Data Aggregator (Go)")

	serials := utils.GenerateSerialNumbers()

	// Batch into groups of 10
	var batches [][]string
	for i := 0; i < len(serials); i += 10 {
		batches = append(batches, serials[i:i+10])
	}

	fmt.Printf("Total Batches: %d\n", len(batches))

	data, err := limiter.ProcessBatches(batches)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Aggregation Complete")
	fmt.Printf("Total Devices Fetched: %d\n", len(data))
}
