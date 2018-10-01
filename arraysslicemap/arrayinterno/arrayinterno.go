//Usando slice para utilizar o mesmo array interno

package main

import (
	"fmt"
)

func main() {

	s1 := make([]int, 10, 20)

	fmt.Println("Slice 1: ", s1)
	s2 := append(s1, 100, 200, 300)
	fmt.Println("Slice 2: ", s2)

	//Será que os slices realmente apontam para o mesmo array interno?
	s1[0] = 500 //Aribuindo 500 ao primeiro elemento do slice s1
	fmt.Println(s1, s2)

	//Resposta: sim! Eles apontam para o mesmo array interno!
}
