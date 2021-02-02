package main

import "fmt"

func modifyArray(a []string) []string {
	a[1] = "x"
	return a
}

func main() {
	array1 := []string{"a", "b", "c"}
	fmt.Printf("The array array1: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array array2: %v\n", array2)
	fmt.Printf("The original array array1: %v\n", array1)
	fmt.Println()
}
