// Arrays and slices
package funcs

import "fmt"

var INTEGER = 3

func Arraysnslices() {
	fmt.Println("Hello World!")
	fmt.Print("\n")

	fmt.Println("Fixed Length:")
	// Fixed-length
	array_one := [...]int{0, 1, 2}
	fmt.Println(array_one)

	// Changing single values
	array_one[2] = 3
	fmt.Println(array_one)

	// Length of the array is fixed
	fmt.Println(len(array_one))

	// Appending will not work
	// array_one = append(array_one, 4, 5)
	// Would throw an error.

	fmt.Println("Variable Length")

	// Variable length
	array_two := []int{10, 11}
	fmt.Println(array_two)
	fmt.Println(len(array_two))

	//Appending DOES work!
	array_two = append(array_two, 13)
	fmt.Println(array_two)

	// Getting a subset of it
	first_element := array_two[0] //golang indexes start at 0
	fmt.Println(first_element)

	second_to_third_element := array_two[1:3] //The 3 indicates that elements up to element 3 are retrieved.
	fmt.Println(second_to_third_element)

	// Two different ways to retrieve exactly the same
	second_element_a := array_two[1]
	second_element_b := array_two[1:2]

	fmt.Println(second_element_a)
	fmt.Println(second_element_b)

	// However, not exactly the same
	// second_element_a is of type int
	// second_element_b is of type int array
	fmt.Printf("Element a of type: %T\n", second_element_a)
	fmt.Printf("Element b of type: %T\n", second_element_b)

	/*
		The following would not work out:
		if second_element_a == second_element_b[] {
			fmt.Println("Heureka!")
		}

		This is because elements of two different types cannot be compared.

		However, the following can be done:
	*/

	if second_element_a == second_element_b[0] {
		fmt.Println("Heureka!")
	}

	// Selecting the first element returns an int which can then be compared.
	fmt.Printf("Element a of type: %T\n", second_element_a)
	fmt.Printf("Subset of element of type: %T\n", second_element_b[0])
}
