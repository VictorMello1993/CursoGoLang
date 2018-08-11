//Funções recursivas

package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Número inválido.")
	case n == 0 || n == 1:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1) //Chamando a função recursiva
		return n * fatorialAnterior, nil
	}
}

func main() {
	num, erro := fatorial(-4)

	if erro != nil {
		fmt.Println(erro)
	} else {
		fmt.Println(num)
	}
}
