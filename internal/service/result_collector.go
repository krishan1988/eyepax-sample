package service

import (
	"context"
	"log"
	"sample/internal/models"
	"sample/internal/store"
	"time"
)

func AggregateResult(ctx context.Context, events []models.Event, resultStore store.AggregatedResultStore) {
	for _, event := range events {
		t, err := time.Parse(time.RFC3339, event.Timestamp)
		if err != nil {
			log.Printf("Unable to process the event. error: %s", err)
			continue
		}
		window := t.Truncate(time.Minute).Format("2006-01-02T15:04:00Z")
		if !resultStore.IsExist(ctx, window) {
			resultStore.Store(ctx, window, models.AggregatedResult{
				WindowStart: window,
			})
		}
		resultStore.Update(ctx, window, event.Value)
	}
}
