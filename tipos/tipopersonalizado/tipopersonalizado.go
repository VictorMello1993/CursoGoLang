//tipos personalizados

package main

import (
	"fmt"
)

//Dando apelido a um tipo primitivo
type nota float64

func (n nota) entre(notaInicio, notaFim float64) bool {
	return float64(n) >= notaInicio && float64(n) <= notaFim
}

func notaParaConceito(n nota) string {
	switch {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 7.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	default:
		return "E"
	}
}

func main() {

	var nt nota //no lugar do float64, o tipo passa a ser nota
	nt = 10
	// fmt.Println(notaParaConceito(10))
	fmt.Println(notaParaConceito(nt))
	fmt.Println(notaParaConceito(5.5))
}
