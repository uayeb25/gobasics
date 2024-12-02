package main

import "fmt"

func increment(x *int) {
	*x = *x + 1
}

func main() {
	fmt.Println("Hello, World!")

	num := 5

	fmt.Println("Value of num before increment is: ", num)
	fmt.Println("Address of num is: ", &num)

	increment(&num)
	increment(&num)
	increment(&num)

	fmt.Println("Value of num after increment is: ", num)

}
