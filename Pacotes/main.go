//Criar modulo 
// go mod init NomeDoModulo 

// Criar um arquivo NomeDoModulo.exe
// go build

// Execultar esse aquivo .exe (ttem todo nosso codigo copilado)
// ./NomeDoModulo

// Oh modulo não atualiza, caso queira atualizar, vc precisa dar aquele msm comando
// go build

// Pegar um pacote da web
// go get github.com/badoux/checkmail

// go mod tidy (vai remover as dependencias no modulo que não estão sendo utilizadas)
package main

import (
	"fmt"
	"modulo/auxiliar"
	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	
	// Pra você referenciar uma função de qualquer pacote
	// você sepre vai usar o nome depois da ultima barra 
	erro := checkmail.ValidateFormat("gabrielcordeirobarroso@gmail.com")
	fmt.Println(erro)
}