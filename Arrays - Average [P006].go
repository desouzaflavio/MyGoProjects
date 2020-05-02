package main

import "fmt"

func main() {

	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Valores:\n", i)
	fmt.Println("Qtd de elementos: ", len(i))

	var avg float64 = mean(i, len(i))

	fmt.Printf("Média: %f", avg)
}

func mean(m []int, size int) float64 {

	var sum int = 0

	for i := 0; i < size; i++ {
		sum += m[i]
	}

	return float64(sum) / float64(size)
}
