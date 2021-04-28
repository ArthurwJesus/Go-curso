//Componentes:Arthur de Jeus,Pedro Giasson e João Paulo
// Linguagem Go lang

// Go não possui classes, mas tem tipos. No caso, tem estruturas. Estruturas são tipos definidos pelo usuário.
//Estruturas (com métodos) são parecidas com classes em outras linguagens.
package main

import "fmt"

//Mitologico é uma struct pq armazena dados
type Mitologico struct {
	Força        int
	Inteligencia int
	Nome         string
	Drop         int
}

func main() {

	fmt.Println("Criaturas existentes nesse bioma:\n")
	Npc := Mitologico{}                          //fazendo um novo "personagem"
	fmt.Printf("NPC: %+v\r\n", Npc)              //vai puxar os dados padrões
	Goblin := Mitologico{18, 2, "Goblinaldo", 2} //adicionando dados
	fmt.Printf("Goblin: %+v\r\n", Goblin)        //imprimindo os dados
	Bruxa := Mitologico{                         //fazendo um novo "monstro" e adicioando dados de outra forma
		Inteligencia: 80,    //adicionando dados
		Força:        70,    //adicionando dados
		Nome:         "ADA", //adicionando dados
		Drop:         8,     //adicionando dados
	}
	fmt.Printf("Bruxa: %+v\r\n", Bruxa) //imprimindo os dados

	Dragao := Mitologico{}                  //fazendo um novo "monstro" e adicioando dados de outra forma
	Dragao.Nome = "Valdohin"                //adicionando dados
	Dragao.Drop = 3                         //adicionando dados
	Dragao.Força = 96                       //adicionando dados
	Dragao.Inteligencia = 70                //adicionando dados
	fmt.Printf("Dragão: %+v\r\n\n", Dragao) //imprimindo os dados

	fmt.Println("Drops:")
	itens := [...]string{"Espada", "Poção", "Escudo", "Grimorio", "Ouro", "Arco"} //[...] fica "ilimitado" ou seja o array não tem um tamanho definido
	for indice, dropadp := range itens {                                          //laço para obter as informações dos drops
		fmt.Printf("Item[%d] -> %s\n\r", indice, dropadp) //vai falar o indice dos drops e qual drop é
	}
}
