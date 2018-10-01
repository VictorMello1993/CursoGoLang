//Função Defer: Função que atrasa a execução de uma setença de código.
//Isso só ocorre quando uma função terminar a execução, antes de retornar algum valor
//É útil para liberar recursos, como fechar um arquivo, fechar conexão com banco de dados, entre
//outros

package main

import (
	"fmt"
)

func obterValorAprovado(num int) int {
	defer fmt.Println("Fim!")

	if num > 5000 {
		fmt.Println("Valor muito alto")
		return 5000
	}

	fmt.Println("Valor muito baixo")
	return num
}

func main() {
	fmt.Println(obterValorAprovado(2000))
	fmt.Println(obterValorAprovado(10000))
}
