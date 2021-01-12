package arrays

func ArrayOfProducts(array []int) []int {
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
