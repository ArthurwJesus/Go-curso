package main

import "fmt"

//Imovel é uma struct pq armazena dados
type Imovel struct {
	x     int
	y     int
	nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("Imovel é: %+v\r\n", casa)
	apartamento := Imovel{18, 29, "Buteco", 30000}
	fmt.Printf("O lugar é: %+v\r\n", apartamento)
	fazenda := Imovel{
		y:     24,
		x:     50,
		nome:  "Fazendo do boi",
		valor: 80000,
	}
	fmt.Printf("A fazendo é: %+v\r\n", fazenda)
	casa.nome = "LAR doce LAR"
	casa.valor = 50000
	casa.x = 20
	casa.y = 20
	fmt.Printf("A casa é: %+v\r\n", casa)
}
