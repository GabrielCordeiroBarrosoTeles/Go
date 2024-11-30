package main

import "fmt"

func main() {
	
	// Aritméticos
	fmt.Println("--------------------")
	soma := 1 + 2
	sub := 1 - 2
	div := 4 / 2
	mult := 5 * 2
	restoDiv := 10 % 2

	fmt.Println(soma,sub,div,mult,restoDiv)

	//ATENÇÃO!!!
	// O Go é fortemente tipado, então não é permitido 
	// fazer nada entre 2 variaveis de tipos diferentes (int8,int16)
	// 	Exemplo de erro
	// 		var numero1 int16 = 10
	// 		var numero2 int32 = 25
	//		soma := numero1 + numero2

	// Atribuição
	fmt.Println("--------------------")
	var variavel1 string = "String 1"
	variavel2 := "String 2 "
	fmt.Println(variavel1,variavel2)

	// Relacionais
	fmt.Println("--------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <=  2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// Operadores Logicos
	fmt.Println("--------------------")
	verdadeiro,falso := true,false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro )

	// Unários
	fmt.Println("--------------------")
	numero := 10

	numero++     //numero = numero + 1
	numero += 15 // numero = numero + 15

	numero--     //numero = numero - 1
	numero -= 20 // numero = numero - 20
	
	numero *= 3 // numero = numero * 3
	numero /= 2 // numero = numero / 2
	numero %= 2 // numero = numero % 2

	fmt.Println(numero)


	// Ternario
	// 	Não existe isso...
	// 		texto := numero > 5 ? "Maior que 5" : "Menor que 5"
    // 	mas a gente pode fazer assim...

	var texto string 
	if numero > 5 {
		texto = "Maior que 5"
	}else{
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}