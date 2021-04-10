//URI Online Judge | 1006 Média 2

package main

import (
	"fmt"
)

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Scanf("%f\n", &a)
	fmt.Scanf("%f\n", &b)
	fmt.Scanf("%f\n", &c)
	media := ((a * 2) + (b * 3) + (c * 5)) / 10
	fmt.Println("MEDIA = %.1f", media)
}

//URI Online Judge | 1001 Extremamente Básico
 package main

import (
    "fmt"
)

func main(){
    var a, b, x int
    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)
    x = a + b
    fmt.Println("X =", x)

}
//ou
package main

 import "fmt"

 func main(){
	var A ,B int
	fmt.Scanf("%d", &A)
	fmt.Scanf("%d",&B)
    fmt.Printf("X = %d\n", A + B)
 }
 //nenhum roda da Time limit exceeded
 
//URI Online Judge | 1050 
DDD
 package main 
 
 import "fmt"

 func main(){
	fmt.Scanf("%d",&DDD)
	if DDD == 61 {
		fmt.Println("Brasilia")
	}
	if DDD == 71{
		fmt.Println("Salvador")
	}
	if DDD == 11 {
		fmt.Println("Sao Paulo")
	}
	if DDD == 21 {
		fmt.Println("Rio de Janeiro")
	}
	if DDD == 32 {
		fmt.Println("Juiz de Fora")
	}
	if DDD == 19{
		fmt.Println("Campinas")
	}
	if DDD ==27 {
		fmt.Println("Vitoria")
	}
	if DDD ==  31{
		fmt.Println("Belo Horizonte")
	}else{
		fmt.Println("DDD não cadastrado")
	}
 }