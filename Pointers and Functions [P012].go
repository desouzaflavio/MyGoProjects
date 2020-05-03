package main

import "fmt"

func main() {

	x := 10

	fmt.Println("Valor de x inicial: ", x)
	zeroing(&x)
	fmt.Println("Valor de x após alteração pela função: ", x)

}

func zeroing(p *int) {
	*p = 0
}
