package tests

import (
	"context"
	"testing"

	"github.com/machester4/greedy-approximation-algorithm/calculator"
)

func TestGreedyCalculator_Calculate(t *testing.T) {
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
	greedySolution := greedyCalculator.Calculate(context.Background())

	// Assert
	if greedySolution.Amount != 9.99 {
		t.Errorf("Expected 9.99, got %f", greedySolution.Amount)
	}
	if len(greedySolution.Items) != 5 {
		t.Errorf("Expected 8 items, got %d", len(greedySolution.Items))
	}
}

// test case with max amount of 500 and items [260, 210, 100, 80, 90]
func TestGreedyCalculator_Calculate_WithMaxAmountOf500(t *testing.T) {
	// Arrange
	items := []calculator.GreedyItem{
		{
			ID:     "1",
			Amount: 260,
		},
		{
			ID:     "2",
			Amount: 210,
		},
		{
			ID:     "3",
			Amount: 100,
		},
		{
			ID:     "4",
			Amount: 90,
		},
		{
			ID:     "5",
			Amount: 80,
		},
	}
	maxAmount := 500.0

	// Act
	greedyCalculator := calculator.NewGreedyCalculator(items, maxAmount)
	greedySolution := greedyCalculator.Calculate(context.Background())

	// Assert
	if greedySolution.Amount != 480 {
		t.Errorf("Expected 480, got %f", greedySolution.Amount)
	}
	if len(greedySolution.Items) != 4 {
		t.Errorf("Expected 4 items, got %d", len(greedySolution.Items))
	}
}
