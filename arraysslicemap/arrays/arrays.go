//Introdução a arrays

package main

import (
	"fmt"
)

func main() {
	//Arrays são estruturas de dados homogêneas (dados de mesmo tipo) e estáticas (fixas)

	//Declarando um array de 3 elementos do tipo float64
	var notas [3]float64
	fmt.Println(notas) //Imprimindo todos os elementos do array com valor zero (valor default para float64)

	//Inicializando elementos de forma indexada
	notas[0], notas[1], notas[2] = 7.5, 8.0, 6.0
	fmt.Println(notas)

	// notas[3] = 10 //Erro de compilação. Índice inválido, pois para um array de 3 elementos, os mesmos são acessados das posições 0 até 2

	total := 0.0
	for i := 0; i < len(notas); i++ {
		total += notas[i]
	}
	media := total / float64(len(notas))
	fmt.Printf("Média: %.2f\n", media)
}
