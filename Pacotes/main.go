//Criar modulo 
// go mod init NomeDoModulo 

// Criar um arquivo NomeDoModulo.exe
// go build

// Execultar esse aquivo .exe (ttem todo nosso codigo copilado)
// ./NomeDoModulo

// Oh modulo não atualiza, caso queira atualizar, vc precisa dar aquele msm comando
// go build
package main

import (
	"modulo/auxiliar"
	"fmt"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
}