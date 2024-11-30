package main

import "fmt"

func main() {
	// Explícito
	var variavel1 string = "Variável 1"

	// Implícito
	variavel2 := "Variável 2"

	fmt.Println(variavel1+ " \n" +variavel2)

	var (
		variavel3 string = "Oi eu sou a variavel 3"
		variavel4 string = "E eu a variavel 4"
	)
	fmt.Println(variavel3+ " \n" +variavel4)

	variavel5,variavel6 := "Variavel 5","Variavel 6"
	fmt.Println(variavel5+ " \n" +variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	// Invertendo os valores das variaveis
	variavel5,variavel6 = variavel6,variavel5
	fmt.Println("Investi os valores das variaveis 5 e 6"+variavel5+ " \n" +variavel6)
	


}