package main

import "fmt"

func main() {

	fmt.Println("Digite um valor:")

	var v int = 0

	fmt.Scanf("%d", &v)
	m, p := half(v)

	fmt.Printf("half: %d\nOdd? %t", m, p)

}

func half(v int) (int, bool) {

	if v%2 == 1 {
		b := true
		return int(v / 2), b
	} else {
		b := false
		return int(v / 2), b
	}
}
