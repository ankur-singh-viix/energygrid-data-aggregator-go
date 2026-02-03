
package limiter

import (
	"fmt"
	"time"
	"energygrid-client-go/internal/api"
)

// Process batches at strict 1 request per second
func ProcessBatches(batches [][]string) ([]api.DeviceData, error) {
	results := []api.DeviceData{}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for i, batch := range batches {
		fmt.Printf(" Processing batch  %d/ %d\n", i+1, len(batches))

		for {
			<-ticker.C

			data, err := api.FetchBatch(batch)
			if err != nil {
				if err == api.ErrRateLimited {
					fmt.Println("Rate limited. Retrying...")
					continue
				}
				return nil, err
			}

			results = append(results, data...)
			break
		}
	}

	return results, nil
}
