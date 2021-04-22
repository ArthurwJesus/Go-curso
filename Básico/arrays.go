//Arrays é fixo

package main

import "fmt"

func main() {
	var teste [3]int
	teste[0] = 10
	teste[1] = 5
	teste[2] = 0
	fmt.Println("Qual a capacidade do Array? ", len(teste)) //ver quantos itens possui
	for indice, valor := range teste {
		fmt.Printf("O valor[%d] é %d\n\r", indice, valor) //vai falar o indice do array e o valor
	}

	cantores := [2]string{"Bon Jovi", "James Arthur"} //com string
	fmt.Printf("Quais cantores há no Array?  \n\r%v\r", cantores)

	paises := [...]string{"Brasil", "Argentina", "Venezulea", "Colombia", "Peru"} //[...] fica "ilimitado" ou seja o array não tem um tamanho definido
	fmt.Println("Qual a capacidade desse Array? ", len(paises))                   //ver quantos itens possui
	for indice, mundi := range paises {                                           //laço para obter as informações dos paises
		fmt.Printf("Pais[%d] é %s\n\r", indice, mundi) //vai falar o indice dos paises e qual pais é
	}

}
