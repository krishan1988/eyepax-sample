package store

import (
	"context"
	"sample/internal/models"
	"sync"
)

type AggregatedResultStore interface {
	Store(ctx context.Context, window string, result models.AggregatedResult)
	IsExist(ctx context.Context, window string) bool
	Update(ctx context.Context, window string, value int)
	GetAll(ctx context.Context) []models.AggregatedResult
}

type aggregatedResultStore struct {
	mu        sync.RWMutex
	windowMap map[string]*models.AggregatedResult
}

func NewAggregatedResultStore() AggregatedResultStore {
	return &aggregatedResultStore{
		windowMap: make(map[string]*models.AggregatedResult),
	}
}

func (a *aggregatedResultStore) Store(_ context.Context, window string, result models.AggregatedResult) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.windowMap[window] = &models.AggregatedResult{
		WindowStart: result.WindowStart,
	}
}
func (a *aggregatedResultStore) IsExist(_ context.Context, window string) bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	_, ok := a.windowMap[window]
	return ok
}
func (a *aggregatedResultStore) Update(_ context.Context, window string, value int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.windowMap[window].EventCount++
	a.windowMap[window].TotalSum += value
}

func (a *aggregatedResultStore) GetAll(_ context.Context) []models.AggregatedResult {
	a.mu.RLock()
	defer a.mu.RUnlock()
	results := make([]models.AggregatedResult, 0)
	for _, val := range a.windowMap {
		results = append(results, *val)
	}
	return results
}
