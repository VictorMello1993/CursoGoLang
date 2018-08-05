//Desafio: Refatorar o exemplo dado na aula de if/else if para executar com switch sem nenhum valor associado

package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {
	switch {
	case n == 10:
		fallthrough
	case n >= 9:
		return "A"
	case n < 9 && n >= 8:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(9.5))
	fmt.Println(notaParaConceito(5.5))
	fmt.Println(notaParaConceito(6))
	fmt.Println(notaParaConceito(4.0))
	fmt.Println(notaParaConceito(2.5))
}
