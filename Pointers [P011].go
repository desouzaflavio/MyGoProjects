package main

import "fmt"

func main() {

	var x int = 5
	var x_pointer *int = &x

	fmt.Println("Valor da Variável x:", x)
	fmt.Println("Endereço da Variável x:", &x)
	fmt.Println("Valor do armazenado no endereço que x_pointer referencia:", *x_pointer)
	fmt.Println("Endereço da Variável x:", x_pointer)
	fmt.Println("Endereço do ponteiro para x:", &x_pointer)

}
