//usado para trabalhar com arryas dinamicos

package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums)) // cap e pra ver a capacidade,len é o comprimento e nums é oq tem nele
	nums = make([]int, 5)                   //inicializa com um comprimento e item de 5 o make faz isso
	fmt.Println(nums, len(nums), cap(nums)) // cap e pra ver a capacidade,len é o comprimento e nums é oq tem nele
	cidades := []string{"Passo Fundo"}      //criou uma slice com o nome passo fundo de tamanho 1 e item 1
	fmt.Println(cidades, len(cidades), cap(cidades))
	cidades = append(cidades, "Ijui")                //append adicionando um novo item dentro do slice cidades
	fmt.Println(cidades, len(cidades), cap(cidades)) // cap e pra ver a capacidade,len é o comprimento e nums é oq tem nele
	nomes := make([]string, 4)                       //criando  e abaixo adicionando
	nomes[0] = "Arthur"
	nomes[1] = "João"
	nomes[2] = "Pedro"
	nomes[3] = "Xaulin"
	fmt.Println("1", nomes, len(nomes), cap(nomes)) // cap e pra ver a capacidade,len é o comprimento e nums é oq tem nele
	nomes[3] = "Cleiton"                            //trocando o nome do item 3
	fmt.Println("2", nomes, len(nomes), cap(nomes)) // cap e pra ver a capacidade,len é o comprimento e nums é oq tem nele
	for indice, nomes := range nomes {              //percorrendo e mostrando o que tem dentro de nomes e a qual lugar está
		fmt.Printf("Nome[%d]= %s\n\r", indice, nomes)
	}
}
