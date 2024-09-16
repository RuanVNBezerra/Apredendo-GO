package main

import (
	"fmt"
)

func main() {
	var nota1 float32
	var nota2 float32
	var nota3 float32

	fmt.Println("Digite a primeira nota: ")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a segunda nota: ")
	fmt.Scanln(&nota2)
	fmt.Println("Digite a terceira nota: ")
	fmt.Scanln(&nota3)

	var media float32

	media = (nota1 + nota2 + nota3) / 3

	if media >= 7 {
		fmt.Println("Parábens, você passou!!!")
	} else if media >= 4 && media < 7 {
		fmt.Println("A sua média está abaixo de 7, por favor, estudar mais!!")
	} else {
		fmt.Println("Você está reprovado por média baixa!")
	}

	fmt.Printf("A média das suas notas são: %.2f", media)

}