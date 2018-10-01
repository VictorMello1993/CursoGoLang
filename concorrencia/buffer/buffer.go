//Channels com buffer

//Diferenças entre a execução de um canal sem buffer e com buffer:

//Quando se utiliza um canal sem buffer, a linha de código onde um dado está sendo enviado ao canal fica
//aguardando até que o dado seja consumido. Caso seja consumido um dado de um canal que ainda não
//recebeu dado, de qualquer maneira, ele espera até o dado chegar nele.

//Por outro lado, quando se utiliza um canal com buffer, a linha de código onde o mesmo recebe dados fica
//aguardando até que o buffer fiquei totalmente cheio, o que gera um bloqueio até que o buffer seja
//reduzido para dar mais espaço para enviar outros dados.

package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3 //Limite do buffer
	ch <- 4
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}

func main() {

	//Criando um canal com buffer de tamanho 3
	c := make(chan int, 3)

	go rotina(c)
	time.Sleep(3 * time.Second) //Com a função sleep, o código consegue enviar mais dados para o canal
	//mesmo quando o buffer fica cheio. Só é possível enviar mais dados
	//quando um dos dados recebidos pelo buffer seja consumido

	fmt.Println(<-c) //Com o 1º uso do buffer, já reduz o espaço para envio de dados. Porém, a linha
	//do println na goroutine não será executada pois o buffer já está cheio

	// fmt.Println(<-c) //Com o uso do buffer pela 2ª vez, o println da goroutine será executado pois já foram liberados
	//2 espaços de buffer
}
