//Operadores ternários. Na realidade, isso não existe em Go, infelizmente. Mas é possível criar alternativas a isso

package main

import (
	"fmt"
)

func obterResultado(nota float64) string {
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(7))
}
