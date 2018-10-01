//Utilizando channels com as goroutines (funcionamento completo da concorrência)

package main

import (
	"fmt"
	"time"
)

//Relembrando: um channel (canal) é a forma de comunicação entre as goroutines, ou seja, é o ponto de sincronismo
//entre as mesmas
//Channel é um tipo como qualquer (int, float, struct interface, etc)

func multiplicar(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base //Enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main() {
	c := make(chan int)
	go multiplicar(2, c) //Com a chamada da goroutine, a função main é executada de maneira assíncrona, sem esperar
	/*o término da execução da goroutine. Ou seja, a goroutine só continuará em execução quando a função main()
	terminar a sua*/

	fmt.Println("A")
	time.Sleep(5 * time.Second)
	//a, b := <-c, <-c /*Recebendo dados do canal proveniente da goutine (aqui é o ponto de sincronismo, a função main
	// aguarda a passagem de valores para um canal
	// da goroutine para depois continuar com a sua
	// execução)*/

	fmt.Println("B")

	// fmt.Println(a, b)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c) //Lendo o 3º valor recebido do canal
	// fmt.Println(<-c) /*ERRO! Deadlock! Todas as goroutines já executaram, e estão paradas. Não existe uma outra
	// goroutine que envie um dado para o canal*/
}
