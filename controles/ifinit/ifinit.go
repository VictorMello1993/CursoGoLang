//Definindo if no bloco de inicialização do código

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano()) //Gerando número aleatório representado por unidade de milisegundos
	r := rand.New(s)
	return r.Intn(10) //Definindo intervalo numérico de 0 até 10
}

func main() {
	if i := numeroAleatorio(); i > 5 { //Definindo if em um bloco de inicialização. Em seguida será comparado se i > 5
		fmt.Println("Ganhou!!!!")
	} else {
		fmt.Println("Perdeu!!!")
	}
}
