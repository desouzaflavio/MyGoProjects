package main

import "fmt"

func main() {

	var fahr float64
	fmt.Println("Enter with fahreint degrees:")
	fmt.Scanf("%f", &fahr)
	conv_FahrToCel(fahr)

}

func conv_FahrToCel(f float64) {

	fmt.Printf("Value in Celsius degrees: %f", float64((f-32)*5/9))
}
