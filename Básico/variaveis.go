// variveis 
package main

import "fmt"
var endereço string = "Ijui" //declaração nome e o tipo = “” -> padrão
var telefone = "0000-0000" // variável já com valor 
var quantidade int//(16,32,64,8-Depende da plataforma cadastrada)// maioria 64  = 0-> padrão
var comprou bool //true e false =false-> padrão
var float64 = 0.00 //(64 ou 32)// maioria 64 tbm = 0.00-> padrão
var palavras rune//unicode

func main(){
    fmt.Println("endereço:  \n", endereço)
    fmt.Println("quantidade: \n", quantidade) 
    fmt.Println("comprou: \n", comprou) 
}

//
package main

import "fmt"

var( // pode ser usadas em todo o codigo 
endereço string = "Ijui" //declaração nome e o tipo = “” -> padrão
telefone = "0000-0000" // variável já com valor 
quantidade int//(16,32,64,8-Depende da plataforma cadastrada)// maioria 64  = 0-> padrão
comprou bool //true e false =false-> padrão
float64 = 0.00 //(64 ou 32)// maioria 64 tbm = 0.00-> padrão
palavras rune//unicode
)


func main(){
    teste :="Valor de teste" // o propio compilador vai achar qual é a varaivel, só disponivel na main agora
    fmt.Println("endereço:  \n", endereço)
    fmt.Println("quantidade: \n", quantidade) 
    fmt.Println("comprou: \n", comprou) 
}

//Go é multiplataforma mas tem que compilar atraves do Go build,roda em qualquer plataforma

//GOOS=windows GOARCH=386 go build -o nomedoaquivo.exe -> Transformar em Wind

//GOOS=linux   GOARCH=arm(rapsberry) go build -o meuapprapsberry -v//(isso ve os pacotes) //https://github.com/ArthurwJesus/Go-curso/blob/main/Básico/funcoesmatematica.go(diretorio de onde esta) 
