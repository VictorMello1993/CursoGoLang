//Operadores lógicos

package main

import (
	"fmt"
)

//Função que retorna 3 valores
func compras(trabalho1, trabalho2 bool) (bool, bool, bool) {
	comprarTv50 := trabalho1 && trabalho2
	comprarTv32 := trabalho1 != trabalho2 //Ou exclusivo
	comprarSorvete := trabalho1 || trabalho2

	return comprarTv50, comprarTv32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := compras(false, false)
	fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
