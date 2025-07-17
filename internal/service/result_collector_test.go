package service

import (
	"context"
	"sample/internal/models"
	"sample/internal/store"
	"testing"
)

func TestAggregateResult(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		Events                  models.Events
		Store                   store.AggregatedResultStore
		ExpectedValuesForWindow map[string]int
	}{
		{
			Events: models.Events{
				{
					Timestamp: "2024-06-06T12:00:00Z",
					Value:     2,
				},
				{
					Timestamp: "2024-06-06T12:00:00Z",
					Value:     2,
				},
			},
			Store: store.NewAggregatedResultStore(),
			ExpectedValuesForWindow: map[string]int{
				"2024-06-06T12:00:00Z": 4,
			},
		},
	}

	for _, test := range tests {
		AggregateResult(ctx, test.Events, test.Store)
		for _, result := range test.Store.GetAll(ctx) {
			val, ok := test.ExpectedValuesForWindow[result.WindowStart]
			if !ok {
				t.Errorf("Expected vale for window %v is incorrect. got: %v, expected: %v", result.WindowStart, result.TotalSum, val)
			}

			if val != result.TotalSum {
				t.Errorf("Expected vale for window %v is incorrect. got: %v, expected: %v", result.WindowStart, result.TotalSum, val)
			}
		}
	}
}
