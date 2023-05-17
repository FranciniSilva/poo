package main

import (
	"fmt"
	"poo/banco/contas"
)

func main() {

	contaExemplo := contas.ContaCorrente{}
	contaExemplo.Depositar(100)

	fmt.Println(contaExemplo.ObterSaldo())

}
