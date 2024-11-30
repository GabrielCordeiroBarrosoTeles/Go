package main

import (
	"errors"
	"fmt"
)

func main() {

	// Int, existem 4 tipos de inteiros
	// 	Só muda a qtd de bits
	//		int8,int16,int32,int64
	// 	se vc colocar apenas "int", ele vai usar a arquitetura do seu PC como base
	// 	Se for 64x o int vai ser "int64"
	var numero1 int16 = -100
	fmt.Println(numero1)

	// unsygned int 
	// Não pode numero com sinal
	var numero2 uint16 = 100
	fmt.Println(numero2)

	// alias - "Apelidos"

	// RUNE = INT32  
	var numero3 rune = 321
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)
	 

	// Float
	// float32,float64
	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123456.78
	fmt.Println(numeroReal2)

	// String
	var str string = "Texto"
	fmt.Println(str)

	// Não existe char, o mais proximo de um char seria assim
	char := 'B'
	// só que ele vai retonar o numero que o B representa na tabela ASCII
	fmt.Println(char) 

	// Boleano
	var booleano bool = true
	fmt.Println(booleano) 

	// Erro
	var erro error = errors.New("Erro interno")
	fmt.Println(erro) 
	




}