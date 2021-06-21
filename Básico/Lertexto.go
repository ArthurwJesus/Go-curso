package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//abre o aruqivo
	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("[lertexto] Houve um Erro 404", err.Error()) //aponta se der erro
		return
	}
	scanner := bufio.NewScanner(arquivo) //pega o arquivo
	for scanner.Scan() {                 //abre ele
		linha := scanner.Text()
		fmt.Println("Cidade:", linha) //mostra os itens
	}
	//outro tipo:
	//leitorCsv := csv.NewReader(arquivo)  //ler o arquivo
	//conteudo, err := leitorCsv.ReadAll() //abre
	//if err != nil {
	//fmt.Println("[lertexto] Houve um Erro 404 com o arquivo csv", err.Error()) //se der erro retorna
	//return
	//}
	//for indiceLinha, linha := range conteudo { //aqui tu ta pedindo pra mostrar o que tem nas linhas
	//fmt.Printf("Linha[%d]: %s\r\n", indiceLinha, linha)
	//for indiceItem, item := range linha {
	//fmt.Printf("Item[%d]: %s\rn", indiceItem, item)
	//}
	//}
	arquivo.Close()
}
