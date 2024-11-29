package auxiliar

import "fmt"

// Função que começa com letra maiúscula podem ser expotadas para outros pacotes
// Função que começa com letra minúscula não pode ser expotadas para outros pacotes,
// apenas no mesmo pacote.
func Escrever (){
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2()
	
}