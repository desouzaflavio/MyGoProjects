package main

import "fmt"

func main() {

	//Declarando Slice de 5 posições para um array de 10 posições
	slice1 := make([]float64, 5, 10)

	//Outra forma de declarar um slice
	arr := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := arr[0:5]

	//3° Forma de declarar slice
	slice3 := []float64{11, 12, 13, 14, 15}

	fmt.Println("Slice 1: ", slice1)
	fmt.Println("Slice 2: ", slice2)
	fmt.Println("Slice 3: ", slice3)

	//Copiando slice
	copy(slice3, slice2)

	fmt.Println("Slice 3: ", slice3)

	//Append para acrecentar um valor ao slice
	slice3 = append(slice3, 7)

	fmt.Println("Slice 3: ", slice3)

}
