package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	contaDaFrancini := ContaCorrente{titular: "Francini", numeroAgencia: 859, numeroConta: 123456, saldo: 500}
	contaDaFrancini2 := ContaCorrente{titular: "Francini", numeroAgencia: 859, numeroConta: 123456, saldo: 500}
	// contaDaCarol := ContaCorrente{titular: "Carol", numeroAgencia: 859, numeroConta: 111222, saldo: 5080}
	// contaDaCarol2 := ContaCorrente{titular: "Carol", numeroAgencia: 859, numeroConta: 111222, saldo: 5080}
	fmt.Println(contaDaFrancini == contaDaFrancini2)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	// fmt.Println(*contaDaCris)

	var contaDaCris2 *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(contaDaCris == contaDaCris2)
	fmt.Println(&contaDaCris)
	fmt.Println(&contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)
}
