//Operadores unários

package main

import "fmt"

func main() {
	x := 1
	y := 2

	//Apenas postfix(pós fixada, depois do operando)
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- //y -= 1 ou y = y - 1
	fmt.Println(y)

	//fmt.Println(x == y--)//ERRO! Em Go, não é permitido utilizar operadores unários em uma comparação, ao contrário de outras linguagens
							//O motivo é devido à sua característica de simplicidade do código
							
}
