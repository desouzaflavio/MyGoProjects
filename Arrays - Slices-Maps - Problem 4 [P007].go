package main

import "fmt"

func main() {

	arr := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17}

	fmt.Println("The Smaller value in Array is: ", findSmaller(arr))
}

func findSmaller(arr []int) int {

	var menor int = arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			menor = arr[i]
		}
	}

	return menor
}
