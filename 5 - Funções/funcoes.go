package main

import "fmt"

// 2 parametros no () e fora do () tem o retorno
func somar (n1 int8,n2 int8) int8 {
	return n1 + n2
}

// Podemos dar mais de 1 retorno em 1 função
func calculosMatematicos (n1,n2 int8) (int8,int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma,sub		
}

func main (){
	soma := somar(10,20)
	fmt.Print(soma)

	// No GO função é um tipo, então vc pode tipar uma variavel com uma func
	var f = func (txt string) string{
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	// Como tem 2 valores sendo retonado, temos que ter 2 variaveis ou...
	resultadoSoma1, resultadoSubtracao1 := calculosMatematicos(10,15)
	fmt.Println(resultadoSoma1, resultadoSubtracao1)

		// Caso vc não queira usar uma variavel é só colocar um underline "_"
		resultadoSoma2, _ := calculosMatematicos(20,15)
		fmt.Println(resultadoSoma2)

		_ , resultadoSubtracao2 := calculosMatematicos(20,15)
		fmt.Println(resultadoSubtracao2)
}