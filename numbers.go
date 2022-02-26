package main

import "fmt"

func fibonacci_numbers(k int) int {
	// Function that takes as input an integer and returns an integer
	// That is, it takes as input number k and returns the k-th fibonacci number
	if k == 0 {
		return 0
	} else if k == 1 {
		return k
	} else {
		return fibonacci_numbers(k-1) + fibonacci_numbers(k-2)
	}
}

func print_fibonacci(k int) {
	// The function fibonacci_numbers only computes the fibonacci number, but does not print it.
	// This function prints it out.
	fmt.Println(fibonacci_numbers(k))
}

func sum(a int, b int) int {
	//Summing two integers
	return a + b
}

func diff(a int, b int) int {
	//Difference between two integers
	return a - b
}

func product(a int, b int) int {
	//Product of two integers
	return a * b
}

func quotient(a float64, b float64) float64 {
	//Quotient of two floats
	//Notice how the two inputs cannot be integers but have to be
	//floats
	return a / b
}

func loop_number(k int, l int) {
	// 1 Sum
	// 2 Product
	// 3 Difference

	for i, j := 0, 0; i < k && j < l; i, j = i+1, j+1 {
		fmt.Println(i, j)
		sum := sum(i, j)
		prod := product(i, j)
		diff := diff(prod, sum)
		fmt.Println(diff)
		fmt.Println("")
	}
}

func loop_numbers(k int, l int) {
	//	The function does the following
	//		Initiate i at 0, if below k, increment by 1
	//		initiate j at 0, if below l, increment by 1
	//		within each loop, do the following:
	//		1: Compute sum of i and j
	//		2: Compute product of i and j
	//		3: Compute difference between the two

	for i := 0; i < k; i = i + 1 {
		for j := 0; j < l; j = j + 1 {
			fmt.Println(i, j)
			sum := sum(i, j)
			prod := product(i, j)
			diff := diff(prod, sum)
			fmt.Println(diff)
			fmt.Println("")
		}
	}
}

func main() {
	//Calls the preceding functions with different numbers
	print_fibonacci(1)
	print_fibonacci(2)
	print_fibonacci(3)
	print_fibonacci(4)
	print_fibonacci(10)

	fmt.Println(sum(12, 4))
	fmt.Println(diff(12, 4))
	fmt.Println(product(12, 4))
	fmt.Println(quotient(12, 4))

	loop_numbers(25, 25)
}
