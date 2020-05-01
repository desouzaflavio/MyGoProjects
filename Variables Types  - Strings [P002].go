package main

import "fmt"

func main() {

	var s string = "Hello World!"
	var s1 string = "Hello"
	var s2 string = "world!"

	//Retorna o tamanho da string
	fmt.Println("Tamanho da String: ", len(s))

	//Retorna o eleento 1 do array Hello World
	fmt.Println("Hello world"[1])
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s1 + " " + s2)

}
