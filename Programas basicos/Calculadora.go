//Crie uma caluladora simples que peça para o usuario digitar os valores, selecionar a operação e que imprima na tela o resultado 
package main

import "fmt"

func main(){
    var A,B,op int
    fmt.Println("_____________________ \n")
    fmt.Println("CALCULADORA BÁSICA")
    fmt.Println("_____________________ \n")
    fmt.Println("Valor de A:")
    fmt.Scanf("%d",&A)
    fmt.Println("Valor de B:")
    fmt.Scanf("%d",&B)
    fmt.Println("_____________________ \n")
    fmt.Println("Operação: \n","1-soma \n","2-subtraçaõ \n","3-Divisão \n","4-Multiplicação")
    fmt.Scanf("%d",&op)
    fmt.Println("_____________________ \n")
    if op == 1{
        fmt.Println("Resultado :", A+B)
        fmt.Println("_____________________ \n")
    }
    if op == 2{
        fmt.Println("Resultado:",A-B)
        fmt.Println("_____________________ \n")
    }
    if op == 3{
        fmt.Println("Resultado:",A/B)
        fmt.Println("_____________________ \n")
    }
    if op ==4{
         fmt.Println("Resultado:",A*B)
         fmt.Println("_____________________ \n")
    }
}
