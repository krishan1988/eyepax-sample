package main

import (
	"context"
	"fmt"
	"log"
	"sample/internal/models"
	"sample/internal/service"
	"sample/internal/store"
)

func main() {
	ctx := context.Background()
	fmt.Println("Time-Based Event Aggregator")

	data := `[
{ "timestamp": "2024-06-06T12:00:00Z", "value": 5 },
{ "timestamp": "2024-06-06T12:00:30Z", "value": 10 },
{ "timestamp": "2024-06-06T12:01:10Z", "value": 3 },
{ "timestamp": "2024-06-06T12:01:50Z", "value": 7 },
{ "timestamp": "2024-06-06T12:02:15Z", "value": 2 }
]`

	events := make(models.Events, 0)
	if err := events.Read(data); err != nil {
		log.Fatalf("Unable to read the event list. error: %s", err)
	}
	resultStore := store.NewAggregatedResultStore()
	service.AggregateResult(ctx, events, resultStore)

	for _, result := range resultStore.GetAll(ctx) {
		fmt.Printf("Window %v | Count %v | Sum %v\n", result.WindowStart, result.EventCount, result.TotalSum)
	}
}
