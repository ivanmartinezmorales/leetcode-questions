package main

func ArrayOfProducts(array []int) []int {
	// Write a function that takes in a non-empty array of integers and returns an array
	// of the same length, where each element in the output array is equal to the product
	// of every other number in the input array

	// In other words, the value at output[i] is equal to the product of every number in the input
	// array other than input[i].

	// Note that you'll be expected to solve this problem without division

	// So we want to return an array that is of size n, where n is the size of the input array

	// When we're on [ 5, 1, 4, 2 ]
	//

	products := []int{}

	for i := range array {
		product := 1
		for j, v := range array {
			if i != j {
				product *= v
			}
		}
		products = append(products, product)
	}
	return products
}
