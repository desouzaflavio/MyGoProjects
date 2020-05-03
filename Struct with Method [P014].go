package main

import (
	"fmt"
)

type Square struct {
	pos_x, pos_y   float64
	weidht, lenght float64
}

func main() {

	c := Square{0, 0, 5, 6}

	fmt.Println("Area do Quadrado:", c.calcArea())
}

func (s *Square) calcArea() float64 {
	return s.weidht * s.lenght
}
