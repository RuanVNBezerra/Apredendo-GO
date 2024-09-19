package main

import (
	"fmt"
)
func main() {
	var inteiro_positivo int

    for { // esse for imita o loop while, enquanto "inteiro_positivo" for maior q 0, sai do loop, se não continua até inserir um numero inteiro
		fmt.Println("Digite um numero inteiro positivo: ")
		fmt.Scanln(&inteiro_positivo)

		if inteiro_positivo > 0 {
			fmt.Println("Numero inteiro inserido!")
            break // momento em que o loop for(while) termina caso o numero for positivo
		} else{
			fmt.Println("Numero inteiro positivo não inserido, tente novamente")
		}
	}

	var soma int
/*
loop para calcular a soma de 1 ate o numero inserido. EX: se o input for 5, será contado 1 + 2 + 3 + 4 + 5 = 15
*/
	for i := 0; i <= inteiro_positivo; i++ { 
        	soma += i
	}

    fmt.Printf("A soma dos numeros de 1 até %d é : %d ", inteiro_positivo, soma)

}