package main

import (
	"fmt"
	"time"
)

func falar(nome, texto string, qte int) {
	for i := 0; i < qte; i++ {
		time.Sleep(time.Second) //Após cada 1 segundo será executada a próxima instrução
		fmt.Printf("%s: %s (iteração %d)\n", nome, texto, i+1)
	}
}

func main() {

	// //Execução sequencial (uma função só será executada depois que a anterior termina)
	// falar("Victor", "Olá, Pedro", 3)
	// falar("Pedro", "Eae, qual é a boa de hoje?", 1)

	//Criando uma goroutine (função que executa de maneira independente de outras funções)
	// go falar("Victor", "Ei...", 500)
	// go falar("Vanderson", "Entendi, parça!", 500)

	//É importante levar em consideração que uma goroutine só será executada quando a função main() terminar
	//a execução.
	go falar("Victor", "Entendi!!!!", 10)
	falar("João", "Parabéns!", 5)

}
