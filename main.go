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
	contaDaCarol := ContaCorrente{titular: "Carol", numeroAgencia: 859, numeroConta: 111222, saldo: 5080}
	fmt.Println(contaDaFrancini, contaDaCarol)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"

	fmt.Println(contaDaCris)
}
