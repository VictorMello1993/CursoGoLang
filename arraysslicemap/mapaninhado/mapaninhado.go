//Maps aninhados

package main

import (
	"fmt"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"V": {
			"Vanderson Guidi": 2500.00,
			"Victor Mello":    2750.00,
		},
		"H": {
			"Humberto Mendonça": 3000.00,
		},
		"D": {
			"Deborah Moreira": 2750.00,
		},
	}

	//Excluindo um dos maps
	delete(funcsPorLetra, "H")

	//Imprimindo todos os maps externos
	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}

	fmt.Println(len(funcsPorLetra))      //Imprimindo o tamanho do map externo
	fmt.Println(len(funcsPorLetra["V"])) //Imprimindo o tamanho do map interno de funcionários que começam com V

	//Percorrendo o map interno de funcionários que começam com V
	for letraV, funcionarios := range funcsPorLetra["V"] {
		fmt.Println(letraV, funcionarios)
	}

	fmt.Printf("\n\n")

	//Usando o for aninhado para percorrer todos os elementos dos maps internos
	for i := range funcsPorLetra {
		fmt.Printf("(%s) %v\n", i, funcsPorLetra[i])
		for j := range funcsPorLetra[i] {
			fmt.Printf(" (%s) %s => %f\n", i, j, funcsPorLetra[i][j])
		}
	}
}
