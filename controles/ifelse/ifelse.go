//Estruturas de controle if/else

package main

import (
	"fmt"
)

func imprimirResultado(nota float64) {
	if nota >= 7 { /*OBS: Em Go, via de regra, não é necessário o uso do parênteses, pois significaria precedência
		Além disso, é obrigatório utilizar blocos de código mesmo que no if possua apenas uma linha para o resultado*/
		fmt.Println("Aprovado com nota", nota)
	} else {
		fmt.Println("Reprovado com nota", nota)
	}
}

func main() {
	imprimirResultado(7.3)
	imprimirResultado(5.1)
}
