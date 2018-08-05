//Estruturas de repetição

package main

import (
	"fmt"
	"time"
)

func main() {

	//Em Go, só existe uma estrutura de repetição, que é o FOR

	//Definindo um FOR de forma semelhante ao While
	i := 1
	for i <= 10 { //Não existe laço while em Go!
		fmt.Println(i)
		i++
	}

	//Definindo FOR tradicional
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	//Misturando estruturas de repetição com condicionais (if dentro de um bloco FOR)
	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("É par")
		} else {
			fmt.Println("É ímpar")
		}
	}

	//Definindo laço infinito
	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second) //Usando Sleep() para evitar que o código rode freneticamente o laço infinito
	}

}
