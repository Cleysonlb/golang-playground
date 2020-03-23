package main

import (
	"estrutura"
	"fmt"
	"matematica"
	"ponteiro"
)

func main() {
	/*
		Funções básicas
	*/
	resultado := matematica.Multiplicadores(3, 10)
	fmt.Printf("Resultado da multiplicação: %d\r\n", resultado)
	resultado = matematica.Soma(3, 10)
	fmt.Printf("Resultado da soma: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Soma, 5, 5)
	fmt.Printf("Resultado da função que chama função de soma: %d\r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 10, 3)
	fmt.Printf("Resultado da divisão: %d\r\n", resultado)
	/*
		#Funções básicas
	*/
	/*
		Funções complexas, retoro de 2 valores da função
	*/
	resultado, resto := matematica.DivisorComResto(12, 5)
	fmt.Printf("Resultado da divisão: %d e o resto foi %d\r\n", resultado, resto)
	/*
		#Funções complexas, retoro de 2 valores da função
	*/
	/*
		Structs
	*/
	estrutura.Main()
	/*
		#Structs
	*/
	/*
		Ponteiro
	*/
	ponteiro.Main()
	/*
		#Ponteiro
	*/
}
