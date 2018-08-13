//Convertendo struct em JSON

package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	//Conversão do struct para JSON
	p1 := produto{1, "Mouse", 180.00, []string{"Promoção", "Eletrônico"}}
	fmt.Println(p1)
	p1ToJSON, _ := json.Marshal(p1) /*Por padrão, o método Marshal recebe como parâmetro um slice de bytes.
	Para que seja retornado no formato de texto, deverá ser realizada uma conversão.*/
	fmt.Println(string(p1ToJSON))

	//Conversão do JSON para struct
	var p2 produto
	jsonToString := `{"id":2,"nome":"Teclado gamer Corsair Strafe RGB", "preco":500.00, "tags":["Promoinfo", "Importado"]}`
	json.Unmarshal([]byte(jsonToString), &p2)
	fmt.Println(p2)
}
