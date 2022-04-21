# Greedy approximation algorithm
This solution was initially intended to solve problems similar to the **knapsack 0/1** or **subset sum** problem.
```
Example: consider the list of numbers [100, 210, 260, 80, 90] 
If the target = 500 then the best solution is 
[100, 210, 80, 90] = 480.
```

## Understanding the algorithm
```
To understand how it works we will use smaller values.
[15, 4, 4, 2, 1, 0.80, 0.50, 0.49] is the list of values.
9.99 is the target.
```
![](/doc/diagram.jpg)

## Using the algorithm
``` go
    // Create a new GreedyApproximation instance
    grapprox := calculator.NewGreedyCalculator(items, maxAmount)

    // Calculate the best solution
    bestSolution := grapprox.Calculate(context.Background())

    // Print the best solution
    fmt.Println(bestSolution)
```


## References
- [Knapsack problem](https://en.wikipedia.org/wiki/Knapsack_problem)
- [Subset sum problem](https://en.wikipedia.org/wiki/Subset_sum_problem)
- [Dynamic programming](https://en.wikipedia.org/wiki/Dynamic_programming)
- [Greedy_algorithm](https://en.wikipedia.org/wiki/Greedy_algorithm)