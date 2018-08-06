//Percorrendo arrays utilizando for range

package main

import (
	"fmt"
)

func main() {

	numeros := [...]int{1, 2, 3, 4} //Compilador irá contar quantos elementos estão definidos num array
	fmt.Println(len(numeros))

	//Percorrendo array com for range (equivalente ao foreach em C#)
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	fmt.Printf("\n")

	//Ignorando o índice utilizando _, para evitar erro de compilação
	for _, num := range numeros {
		fmt.Println(num)
	}

}
