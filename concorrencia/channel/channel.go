//Introdução aos channels (canais)
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1) //Criando um canal de inteiros, e o 2º parâmetro é o tamanho do buffer

	ch <- 1 //enviando dados para o canal (escrita)
	<-ch    //recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch) //Imprimindo o valor lido do canal
}
