package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4}

	//passando valores do slice como parametro
	fmt.Println(Sum(sl...))

	//Adicionando parametros manualmente
	fmt.Println(Sum(1, 2, 3, 4))
}

func Sum(args ...int) int {
	t := 0
	for _, v := range args {
		t += v
	}
	return t
}
