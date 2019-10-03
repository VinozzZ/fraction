# multiplyFractions

A method for multiply fractions.

## How to run the program

- In the root folder
- `go test` will run test in `fraction_test.go`

## Improvment

- how to improve the speed of multiplication with parallel execution?

    1. We can run the multiplication for numerator and denominator in separate goroutine so these two operation can run in parallel.
    2. if the inputs are large numbers, we can also split up each number by bytes so that we can do multiplication separately. And then, we will be able to add up each result to get the final number.
    
