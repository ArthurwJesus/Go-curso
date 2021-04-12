//Ponteiros apontam para o endereço da memória, servem para:
//Alocação dinâmica de memória
//Manipulação de arrays.
//Para retornar mais de um valor em uma função.
//Referência para listas, pilhas, árvores e grafos.
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
	casa := new(Imovel) //criando um ponteiro
	fmt.Printf("Casa é: %p - %+v\r\n",&casa,casa) //vai apontar para onde está o endereço
	chacara:= Imovel{12,28,"Boi Boi", 289999}
	fmt.Printf("Chacara é: %p - %+v\r\n",&chacara,chacara)
}
//mudando valores e mxendo com o ponteiro
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
	casa := new(Imovel) //criando um ponteiro
	fmt.Printf("Casa é: %p - %+v\r\n",&casa,casa) //vai apontar para onde está o endereço
	fazendinha:= Imovel{12,28,"Boi Boi", 289999}
	fmt.Printf("Fazenda é: %p - %+v\r\n",&fazendinha,fazendinha)
	mudaImovel(&fazendinha) //vai imprimir a mudança
	fmt.Printf("Chacara é: %p - %+v\r\n",&fazendinha,fazendinha)
}
func mudaImovel(imovel * Imovel){
	imovel.x = imovel.x +11
	imovel.y = imovel.y - 4
}
