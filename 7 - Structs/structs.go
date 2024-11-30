package main

import "fmt"

// Structs: é uma coleção de campos
type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct{
	logradouro string
	numero uint8
}

func main() {

	var usuario1 usuario
	usuario1.nome = "Gabriel Cordeiro"
	usuario1.idade = 19
	fmt.Println(usuario1)

	enderecoExemplo := endereco{"Rua D" , 10}

	usuario2 := usuario{"Gabriel",19,enderecoExemplo} 
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Gabriel"}
	fmt.Println(usuario3)

	usuario4 := usuario{idade:19}
	fmt.Println(usuario4)

}