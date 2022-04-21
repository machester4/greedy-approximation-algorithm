package tests

import (
	"context"
	"testing"

	"github.com/machester4/greedy-approximation-algorithm/calculator"
)

func BenchmarkGreedyCalculator_Calculate(b *testing.B) {
	// Arrange
	items := []calculator.GreedyItem{
		{
			ID:     "1",
			Amount: 10,
		},
		{
			ID:     "2",
			Amount: 4,
		},
		{
			ID:     "3",
			Amount: 4,
		},
		{
			ID:     "4",
			Amount: 2,
		},
		{
			ID:     "5",
			Amount: 1,
		},
		{
			ID:     "6",
			Amount: 0.80,
		},
		{
			ID:     "7",
			Amount: 0.50,
		},
		{
			ID:     "8",
			Amount: 0.49,
		},
	}
	maxAmount := 9.99

	// Act
	greedyCalculator := calculator.NewGreedyCalculator(items, maxAmount)

	// Benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		greedyCalculator.Calculate(context.Background())
	}
}
