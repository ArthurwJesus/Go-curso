//Mapas ou Cache armazenam dados para que possam ser atendidos mais rapidamente em solicitações futuras.
package main

import "fmt"

type Imovel struct {
	X     int
	Y     int
	Nome  string
	Valor int
}

var cache map[string]Imovel

func main() {
	cache = make(map[string]Imovel, 0)
	apartamento := Imovel{18, 29, "Buteco", 30000}

	casa := Imovel{}
	casa.Nome = "Casa Verde"
	casa.X = 18
	casa.Y = 25
	casa.Valor = 6000

	cache["Casa Verde"] = casa

	fmt.Println("Há uma casa verde no chace?")
	Imovel, achou := cache["Casa Verde"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", Imovel)
	}

	cache[apartamento.Nome] = apartamento             //pode achar algo so pelo nome
	fmt.Println("Quantos item no cache?", len(cache)) //pra ver quantos itens tem

	for chave, Imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, Imovel) //aqui ele vai ler os imoveis na tela
	}
	Imovel, achou = cache["Casa Verde"] //excluir a casa verde
	if achou {                          //se achou
		delete(cache, Imovel.Nome) //exclui
	}
	fmt.Println("___________________________________________________________")
	fmt.Println("Quantos item no cache depois da exclusão?", len(cache)) //pra ver quantos itens tem
	for chave, Imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, Imovel) //aqui ele vai ler os imoveis na tela novamente
	}
}
