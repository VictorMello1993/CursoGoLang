//Structs aninhadas

package main

import (
	"fmt"
)

//Criando um item de pedido do produto
type itemPedido struct {
	produtoID  int
	nome       string
	quantidade int
	preco      float64
}

//Criando um pedido com histórico de itens contendo informações a respeito do produto
type pedido struct {
	userID int
	itens  []itemPedido
}

//Calcular o preço total do pedido
func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.quantidade)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 554654,
		itens: []itemPedido{
			itemPedido{produtoID: 11111, nome: "Monitor Dell", quantidade: 1, preco: 1350.00},
			itemPedido{produtoID: 11112, nome: "GeForce GTX 1060 6 GB", quantidade: 1, preco: 1400},
			itemPedido{produtoID: 11113, nome: "Processador Intel i5 7500", quantidade: 1, preco: 850.00},
			itemPedido{produtoID: 11114, nome: "Pente de memória RAM 8 GB HyperX", quantidade: 2, preco: 450.00},
			itemPedido{produtoID: 11115, nome: "Teclado Corsair Strafe RGB", quantidade: 1, preco: 400.00},
			itemPedido{produtoID: 11116, nome: "Mouse Razer DeathAdder Chroma", quantidade: 1, preco: 500.00},
			itemPedido{produtoID: 11117, nome: "Mouse pad HyperX Speed médio", quantidade: 1, preco: 100.00},
			itemPedido{produtoID: 11118, nome: "Fonte Thermaltake 700W Selo 80 Plus White", quantidade: 1, preco: 400.00},
			itemPedido{produtoID: 11119, nome: "Placa mãe Asus Prime B250M Plus BR LGA 1100", quantidade: 1, preco: 500.00},
			itemPedido{produtoID: 20000, nome: "Gabinete H27 Thermaltake Mid Tower", quantidade: 1, preco: 250.00},
		},
	}
	fmt.Printf("Valor total: R$ %.2f", pedido.valorTotal())
}
