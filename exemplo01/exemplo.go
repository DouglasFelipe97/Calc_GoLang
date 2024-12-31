package main

import "fmt"

func main() {
	var validacao int = 1
	for validacao != 0 {
		/*VARIAVEIS LOCAIS*/
		var numero1 int = 0
		var numero2 int = 0
		var escolha int = 0
		var resultado float32 = 0
		var modulo int = 0
		var operacao string
		/*================*/
		for numero1 == 0 {
			fmt.Print("Digite o primeiro número: ")
			fmt.Scanf("%d\n", &numero1)
			if numero1 == 0 {
				fmt.Print("Escolha um numero diferente de 0!\n\n")
			}
		}
		for numero2 == 0 {
			fmt.Print("Digite o segundo número: ")
			fmt.Scanf("%d\n", &numero2)
			if numero2 == 0 {
				fmt.Print("Escolha um numero diferente de 0!\n\n")
			}
		}
		fmt.Print("\nEscolha a operação: \n1-Soma.\n2-Subtração.\n3-Divisão\n4-Multiplicação\n5-Modulo\n0-Sair\nEscolha: ")
		fmt.Scanf("%d\n", &escolha)
		switch escolha {
		case 0:
			fmt.Print("Fechando a calculadora!")
			validacao = 0
			break
		case 1:
			resultado = float32(numero1) + float32(numero2)
			operacao = "soma"
			break
		case 2:
			resultado = float32(numero1) - float32(numero2)
			operacao = "subtração"
			break
		case 3:
			resultado = float32(numero1) / float32(numero2)
			operacao = "divisão"
			break
		case 4:
			resultado = float32(numero1) * float32(numero2)
			operacao = "multiplicação"
			break
		case 5:
			modulo = (numero1) % (numero2)
			operacao = "modulo"
			break
		default:
			fmt.Print("\nOperação inválida!\n")
		}
		if escolha >= 1 && escolha <= 5 {
			if escolha != 5 {
				fmt.Printf("O resultado da %s entre %d e %d é %.2f\n\n\n\n", operacao, numero1, numero2, resultado)
			} else {
				fmt.Printf("O resultado da %s entre %d e %d é %d\n\n\n\n", operacao, numero1, numero2, modulo)
			}

		}
	}
}
