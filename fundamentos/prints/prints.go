//Comandos de impressão de dados (prints)

package main

import (
	"fmt"
)

func main() {
	/*Ao executar, uma vez declarado o print em múltiplas linhas,
	não serão quebradas as linhas em tempo de execução*/
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	//Neste caso, com Println, a cada chamada, será impressa uma string com quebra de linha
	fmt.Println(" Nova")
	fmt.Println("linha")

	x := 3.141516

	// fmt.Println("O valor de x é " + x) /*ERRO! Não será permitida a concatenação de uma string com o valor de tipo diferente*/

}
