package main

import (
	"fmt"
)

var numero_positivo int

func main() {
// 1. Solicite ao usuário para inserir um número inteiro positivo n.
 for{
	fmt.Println("Digite um número inteiro positivo: ")
	fmt.Scanln(&numero_positivo)
    // 2. Verifique se o número inserido é positivo. Caso contrário, solicite novamente até que um número válido seja fornecido.
    if numero_positivo > 0 {
		fmt.Println("O número inserido é positivo!")
		break
	}else {
		fmt.Println("Oque foi inserido, não condiz com oque pede!")
	}
 }
// 3. Use um loop para imprimir todos os números pares de 0 até n.
  contador := 0
  
  fmt.Printf("Numeros pares de 0 até: %d",numero_positivo)

  for i := 0; i <= numero_positivo; i++ {
	if i % 2 == 0 {
		fmt.Printf("%d", i)
		contador ++
	}
  }
  fmt.Println()
// 4. Ao final, imprima quantos números pares foram exibidos.
 
fmt.Printf("A quantidade de números pares é: %d\n", contador)
}