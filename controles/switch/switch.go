//Comando switch Exemplo 1

package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough //Este comando permite executar o próximo case. Na linguagem Go, não existe o método tradicional de por break em cada case
	case 9:
		return "A"
	case 8, 7: //Comparando múltiplos valores em um case
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(5))
	fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(2))
	fmt.Println(notaParaConceito(7.5))
	fmt.Println(notaParaConceito(-4))
}
