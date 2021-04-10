# shipping-box-problem
Go solution to verify the best box size that fits a list of products, concurrent and non-concurrent way to check the best approach in terms of performance

# Execution
Just run the tests using `go test -bench=.` and see the results

# Improvements
In this versions, the concurrent way is worse than the non-concurrent.
I guess the reason is that I need to order the array of results after all the executions.
To improve that, perhaps I need to force the routines to run in a specific order and stop
the executions after finding the first box.
I'll check that later, if you want to help, fell free to send your pull request.
