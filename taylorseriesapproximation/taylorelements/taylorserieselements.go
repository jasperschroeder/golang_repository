package taylorelements 

import (
	"math"
)

func factorial(n int) int {
	if int(n) <= 1 {
		return 1 
	}
	return int(n) * factorial(int(n)-1)
}

func ExponentialElements(x float64, n int) float64 {
	return math.Pow(x, float64(n)) / float64(factorial(n))
}

func SinusElements(x float64, n int) float64 {
	return math.Pow(-1, float64(n)) * math.Pow(x, float64(2*n+1)) / float64(factorial(2*n+1))
}

func CosineElements(x float64, n int) float64 {
	return math.Pow(-1, float64(n)) * math.Pow(x, float64(2*n)) / float64(factorial(2*n))
}

//f(x) = ln(1+x) 
func LnofoneplusxElements(x float64, n int) float64 {
	if n == 0 {
		return 0 
	}
	return math.Pow(-1, float64(n+1)) * math.Pow(x, float64(n)) / float64(n)
}

func TaninverseElements(x float64, n int) float64 {
	return math.Pow(-1, float64(n)) * math.Pow(x, float64(2*n+1)) / float64(2*n+1)
}

