package main

import "fmt"

func main() {
    for i := 0; i <10; i++{ // laço que vai percorer do 0 ao 9
    fmt.Println("Qual o valor de I?",i) //vai impimir os valores na tela
    }
// agora como posso fazer um while?
    valor := 0 
    teste:=true
    for teste{//faz enquanto o teste for igual a verdadeiro
        valor ++ //vai incrementar
        if valor%5 ==0{ //se o valor divido e o resto da divisão for 5 é igual a 0 
            teste = false //se não for é falso
        }
        fmt.Println("Numero é: ",valor)
    }
}
//Uri URI Online Judge | 1114 -Senha Fixa usando o break que é pra parar

package main

import "fmt"

func main() {
    var( //declara as variaveis
        senha int
        teste bool
        )
    senha = 0 //inicia o valor da senha com 0
    teste = true //pra testar
    fmt.Println("Qual a senha:")//perguntar na tela
    for teste{ //se o teste for verdadeiro ele testa
        fmt.Scanf("%d",&senha)//recebe a senha
        if senha !=2002{ //se for diferente de 2002 ele é falso, fecha o programa e aparece na tela Senha invalida
            fmt.Println("Senha invalida")
            teste = false
        }else{ //senão acesso permitido 
            fmt.Println("Acesso permitido")
            break //Precisa do break senão ele continua pedindo e vira infinito 
        }
    }
}
//Range for
//o Golang só apresenta o for e assim como o if ele não necessita de ()
package main

import "fmt"

func main() {
    
    texto:="Eu adoro escrever programas usando GO"
    for indice, letra:= range texto{ //equivalente ao range for,retornar uma lista de números inteiros.
        fmt.Println("Texto [%d]= %q\r\n",indice,letra)
    }
}