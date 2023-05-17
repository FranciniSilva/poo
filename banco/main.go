package main

import (
	"fmt"
	"poo/banco/contas"
)

func PagarVoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarVoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaPati := contas.ContaCorrente{}
	contaDaPati.Depositar(600)
	PagarVoleto(&contaDaPati, 60)
	fmt.Println(contaDaPati.ObterSaldo())

}
