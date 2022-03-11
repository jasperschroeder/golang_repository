package main

import (
	"fmt"
	"funcs/funcs"
)

func main() {
	fmt.Println("Hello World!")

	// Running function with imported variables from file
	fmt.Println("Hello, my name is", Name, ".")
	fmt.Println("I live in", City, ".")

	// Running function with imported variables from file in different local package
	fmt.Println("Let's import and print out number 3:")
	fmt.Println(funcs.INTEGER)

	// Importing and running functions from a package
	funcs.Arraysnslices()
	funcs.All_funcs()

}
