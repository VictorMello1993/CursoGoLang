//Declaração de uma função em Go

package main

import "fmt"

//Função que retorna um valor de um determinado tipo
func somar(a int, b int) int {
	return a + b
}

//Função que não retorna nenhum valor
func imprimir(valor int) {
	fmt.Println(valor)
}

// func main() {
// 	resultado := somar(3, 4)
// 	imprimir(resultado)
// }
