package main

import (
	"fmt"
	"runtime"
)

func main() {

	//Determinando o número de processadores (núcleos) físicos existentes no sistema
	fmt.Println(runtime.NumCPU())
}
