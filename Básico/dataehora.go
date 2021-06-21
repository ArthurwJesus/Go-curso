//Golang chama seus pacotes de módulos e vice-versa
package main

import (
	"os"   //importa o pacote da criação do txt
	"time" //importa o pacote data e hora
)

func main() {
	file, err := os.Create("Data e hora.txt") //cria o arquivo
	if err != nil {                           //se der problema ele retorna
		return
	}
	defer file.Close()                                       //fecha o prgrama
	data := time.Now()                                       //pega a data e hora atual da maquina
	file.WriteString(data.Format(("02/Jan/2006 15:04:05 "))) // aqui vai escrever no txt e dessa forma ela fica "bonito" -> 10/Apr/2021 19:56:58 //a parte da data ali dentro do codigo não deve ser mudada pois é um padrão Go de qualidade
}
