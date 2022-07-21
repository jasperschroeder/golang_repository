package main 

import (
	"fmt"
	"math"
	taylorelements "github.com/jasperschroeder/golang_repository/taylorseriesapproximation/taylorelements"
)

func taylorseriesapproximation(f func(float64, int) float64, x float64, j int) float64 {
	ch := make(chan float64) 
	for k := 0; k < j; k++ {
		go derivative(ch, f, x, k)
	}
	approx := 0.0
	for k := 0; k < j; k++ {
		approx += <- ch 
	}
	return approx
}


func derivative(ch chan float64, f func(float64, int) float64, x float64, j int) {
	ch <-  f(x,j)
}

func main () {
	fmt.Println("In this module, I will make use of goroutines and channels to compute taylor series expansions.\n")

	numbers := [5]int{2,3,4,5,6}

	//Exponential 
	fmt.Println("Starting with exponential function, evaluated at x=5 with expansion of order 3, 5, 10, 20, 25:")
	fmt.Printf("True value: %f\n", math.Exp(float64(5)))
	for _, number := range [5]int{3,5,10,20,25} {
		fmt.Printf("Approximation at order %g: %g\n", float64(number), taylorseriesapproximation(taylorelements.ExponentialElements, 5, number))
	}

	//Sinus
	fmt.Println("\nSinus function, evaluated at x=0.1 with expansion of order 2, 3, 4, 5, 6:")
	fmt.Printf("True value: %f\n", math.Sin(0.1))
	for _, number := range numbers {
		fmt.Printf("Approximation at order %g: %g\n", float64(number), taylorseriesapproximation(taylorelements.SinusElements, 0.1, number))
	}
	
	//Cosine 
	fmt.Println("\nCosine function, evaluated at x=0.15 with expansion of order 2, 3, 4, 5, 6:")
	fmt.Printf("True value: %f\n", math.Cos(0.15))
	for _, number := range numbers {
		fmt.Printf("Approximation at order %g: %g\n", float64(number), taylorseriesapproximation(taylorelements.CosineElements, 0.15, number))
	}
	
	//Natural log of 1 + x
	fmt.Println("\nNatural logarithm of x + 1, evaluated at x=0.5 with expansion of order 2, 3, 4, 5, 6:")
	fmt.Printf("True value: %f\n", math.Log(1+0.5))
	for _, number := range numbers {
		fmt.Printf("Approximation at order %g: %g\n", float64(number), taylorseriesapproximation(taylorelements.LnofoneplusxElements, 0.5, number))
	}

	// Taninverse 
	fmt.Println("\nTaninverse, evaluated at x = 0.35 with expansion of order 2, 3, 4, 5, 6:")
	fmt.Printf("True value: %f\n", float64(math.Atan(0.35)))
	for _, number := range numbers {
		fmt.Printf("Approximation at order %g: %g\n", float64(number), taylorseriesapproximation(taylorelements.TaninverseElements, 0.35, number))
	}
}