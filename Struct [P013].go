package main

import (
	"fmt"
	"math"
)

func main() {

	c := Circle{10, 5, 7}

	c.x = 10
	c.y = 5
	c.r = 7

	fmt.Println(c.x)
	fmt.Println(c.y)
	fmt.Println("area do circulo", calcCircleArea(&c))

}

type Circle struct {
	x, y, r float64
}

func calcCircleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}
