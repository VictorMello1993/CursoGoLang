//Comando switch exemplo 2

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch { //switch true. Com switch sem nenhum valor associado, ele busca um case cujo resultado é true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}
