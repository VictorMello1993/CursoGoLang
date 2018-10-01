//Inicialização dos maps

package main

import (
	"fmt"
)

func main() {
	//Declarando e inicializando um map de forma literal
	funcsESalarios := map[string]float64{
		"Victor Mello":    5000.00,
		"Vanderson Guidi": 1250.50,
		"Camilla Batista": 4500.00,
		"Anderson Lopes":  3750.50,
	}

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
	fmt.Printf("\n")

	//Ignorando as chaves, imprimindo apenas os valores
	for _, salario := range funcsESalarios {
		fmt.Println(salario)
	}
}
