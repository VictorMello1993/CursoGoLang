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

	//fmt.Println("O valor de x é " + x) /*ERRO! Não será permitida a concatenação de uma
	//string com o valor de tipo diferente*/

	//Algumas forma de resolver o problema relatado acima:
	//1- Através da função Sprint
	xs := fmt.Sprint(x) //Em vez retornar uma string no console, a função Sprint retorna uma versão de x no formato de string
	fmt.Println("O valor de x é", xs)

	//2- Imprindo o Println utilizando concatenação de strings
	fmt.Println("O valor de x é " + xs)

	//3- Em vez de criar uma nova variável de um mesmo tipo que será impresso no console, utilizar Println sem usar
	//concatenação
	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor de x é %.2f", x) //Imprimindo o float com 2 casas decimais utilizando Printf

	a := 1
	b := 1.999999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d) //OBS: modificador %t é utilizado para imprimir valores do tipo bool
	fmt.Printf("\n%v %v %v %v", a, b, c, d)         //O modificador %v é utilizado para imprimir variáveis de qualquer tipo

	

}
