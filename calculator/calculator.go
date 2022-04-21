package calculator

import (
	"context"
	"sort"
)

// greedy calculator phases
// 1. fill the bag until an item exceeds the maximum amount
// 2. look for a reconciliation between the products in the bag and the products outside it,
// 	  exchanging one with another until reaching a solution that does not exceed the maximum total.
// 	2.1. if the solution does not exceed the maximum total, continue with phase 1.
// 	2.2. if the solution exceeds the maximum total, return best solution so far.

type GreedySolution struct {
	Items  []GreedyItem
	Amount float64
}

type GreedyItem struct {
	ID     string
	Amount float64
}

type greedyCalculator struct {
	items             []GreedyItem
	maxAmount         float64
	bestSolutionSoFar GreedySolution
	currentSolution   GreedySolution
}

func (g *greedyCalculator) Calculate(ctx context.Context) GreedySolution {
	// if amount of best solution so far is equal to the maximum amount,
	// should end recursion and return best solution so far
	if g.bestSolutionSoFar.Amount == g.maxAmount {
		return g.bestSolutionSoFar
	}

	// if has conflict in current solution, should try to reconciliate before adding new item
	// if can't reconciliate, should end recursion and return best solution so far
	if g.currentSolution.Amount > g.maxAmount {
		hasReconciliation := g.reconciliate(ctx)
		if !hasReconciliation {
			return g.bestSolutionSoFar
		}
	} else if g.currentSolution.Amount > g.bestSolutionSoFar.Amount {
		// if current solution is better than the best solution so far, update the best solution so far
		g.bestSolutionSoFar = g.currentSolution
	}

	// if items are empty, should end recursion and return best solution so far
	if len(g.items) == 0 {
		return g.bestSolutionSoFar
	}

	// get current item and remove it from the list
	item := g.items[0]
	g.items = g.items[1:]

	// if item amount is greater than the maximum amount, skip it
	if item.Amount > g.maxAmount {
		return g.Calculate(ctx)
	}

	// Add item to current solution
	g.currentSolution.Items = append(g.currentSolution.Items, item)
	g.currentSolution.Amount = g.currentSolution.Amount + item.Amount

	return g.Calculate(ctx)
}

func (g *greedyCalculator) reconciliate(ctx context.Context) bool {
	cf := newConflictMediator(g.items, g.currentSolution, g.maxAmount)
	resolution := cf.Resolve(ctx)

	if resolution.solution.Amount > g.bestSolutionSoFar.Amount {
		g.bestSolutionSoFar = resolution.solution
		g.currentSolution = resolution.solution

		// remove items used in resolution from the items list
		g.items = g.items[resolution.index+1:]

		return true
	}

	return false
}

type ConflictMediator struct {
	items          []GreedyItem
	conflict       GreedySolution
	maxAmount      float64
	solutionsTable [][]GreedySolution
}

type conflictMediatorOutput struct {
	solution GreedySolution
	index    int
}

func (cm *ConflictMediator) Resolve(ctx context.Context) conflictMediatorOutput {
	cm.solutionsTable = make([][]GreedySolution, len(cm.conflict.Items))
	bestResolution := GreedySolution{}
	bestResolutionXDepth := 0

	for y := 0; y < len(cm.conflict.Items); y++ {
		cm.solutionsTable[y] = make([]GreedySolution, len(cm.items))

		for x := 0; x < len(cm.items); x++ {
			s := cm.buildSolution(ctx, x, y)

			if s.Amount == cm.maxAmount {
				return conflictMediatorOutput{
					solution: s,
					index:    x,
				}
			}

			if s.Amount > cm.maxAmount {
				continue
			}

			if s.Amount > bestResolution.Amount {
				bestResolution = s
				bestResolutionXDepth = x
			}

			cm.solutionsTable[y][x] = s
		}
	}

	return conflictMediatorOutput{
		solution: bestResolution,
		index:    bestResolutionXDepth,
	}
}

func (ct ConflictMediator) buildSolution(ctx context.Context, x, y int) GreedySolution {
	if x == 0 {
		s := GreedySolution{
			Items:  exclude(ct.conflict.Items, y),
			Amount: ct.conflict.Amount - ct.conflict.Items[y].Amount,
		}

		s.Items = append(s.Items, ct.items[x])
		s.Amount = s.Amount + ct.items[x].Amount

		return s
	}

	lastSolution := ct.solutionsTable[y][x-1]
	s := GreedySolution{
		Items:  append(lastSolution.Items, ct.items[x]),
		Amount: lastSolution.Amount + ct.items[x].Amount,
	}

	return s
}

func exclude(slice []GreedyItem, index int) []GreedyItem {
	sl := make([]GreedyItem, len(slice)-1)
	copy(sl, slice[:index])

	return sl
}

func newConflictMediator(items []GreedyItem, conflict GreedySolution, maxAmount float64) *ConflictMediator {
	return &ConflictMediator{
		items:     items,
		conflict:  conflict,
		maxAmount: maxAmount,
	}
}

func NewGreedyCalculator(items []GreedyItem, maxAmount float64) *greedyCalculator {
	// sort items by amount in descending order
	sort.Slice(items, func(i, j int) bool {
		return items[i].Amount > items[j].Amount
	})

	return &greedyCalculator{
		items:             items,
		maxAmount:         maxAmount,
		bestSolutionSoFar: GreedySolution{},
		currentSolution:   GreedySolution{},
	}
}
