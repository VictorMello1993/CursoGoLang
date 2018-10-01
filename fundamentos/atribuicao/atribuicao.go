//Operadores de atribuição
package main

import (
	"fmt"
)

func main() {
	//Atribuição simples
	var b byte = 3
	fmt.Println(b)

	//Operadores de atribuição de incremento
	i := 3 //Inferência de tipo
	i += 3 //i = i + 3
	i -= 3 //i = i - 3
	i /= 2 //i = i / 2
	i *= 2 //i = i * 2
	i %= 2 //i = i % 2

	fmt.Println(i)

	//Atribuição com múltiplos valores
	x, y := 1, 2
	fmt.Println(x, y)

	//Substituição de valores entre variáveis
	x, y = y, x //Antes x = 1, y =2
	//Depois x = 2, y = 1. Isso elimina a necessidade de criar uma variável auxiliar para substituir
	fmt.Println(x, y)

}
