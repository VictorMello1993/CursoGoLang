//Encontrando números primos utilizando channels

package main

import (
	"fmt"
	"time"
)

func ehPrimo(num int) bool {
	divisores := 0
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			divisores++
		}
	}

	if divisores == 2 {
		return true
	}

	return false
}

func primos(n int, c chan int) {
	inicio := 2
	for i := 0; i < n; i++ {
		for primo := inicio; ; primo++ {
			if ehPrimo(primo) {
				c <- primo
				inicio = primo + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c)
}

func main() {
	ch := make(chan int, 30) //Canal com buffer de tamanho 30 para representar a quantidade de números primos a serem armazenados
	go primos(cap(ch), ch)
	for primo := range ch {
		fmt.Printf("%d ", primo)
	}
}
