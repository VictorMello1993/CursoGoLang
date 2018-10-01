package main

import (
	"fmt"
)

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	// fmt.Printf("%v %v %v %v %v", a, b, c, d, e) //Imprimindo os valores default para cada tipo de variável
	fmt.Printf("%v %v %v %q %v", a, b, c, d, e) //O modificador %q mostra a string vazia no console
}
