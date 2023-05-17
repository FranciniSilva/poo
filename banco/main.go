package main

import (
	"fmt"
	"poo/banco/clientes"
	"poo/banco/contas"
)

func main() {

	clienteBruno := clientes.Titular{
		Nome:      "Bruno",
		CPF:       "123.111.123.12",
		Profissao: "Desenvolvedor",
	}

	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 845, NumeroConta: 123456, Saldo: 300}

	fmt.Println(contaDoBruno)

}
