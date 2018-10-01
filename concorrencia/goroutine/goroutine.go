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

	//É importante levar em consideração que uma goroutine só será executada quando a função main() estiver funcionando
	//Porém, por sua natureza, como uma goroutine é executada de maneira independente, ela desconhece a execução
	//da função main (ou uma thread principal. Por analogia, uma goroutine funciona como uma thread separada), e
	//a mesma termina antes da execução das goroutines. Então, para que a função main espere que as outras funções
	//independentes terminem a execução, utilizam-se os channels para realizar a comunicação entre elas
	go falar("Victor", "Entendi!!!!", 10)
	falar("João", "Parabéns!", 5)

}
