package main

import (
	"fmt"
	"github.com/ygorsimoes/LearningGo/ByteBank/clientes"
	"github.com/ygorsimoes/LearningGo/ByteBank/contas"
)

func main() {
	contaDoJoao := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "João Santos",
		CPF: "45.234.213-67",
		Profissao: "Programador"},
		NumeroAgencia: 12345,
		NumeroConta: 22,
		Saldo: 100}

	contaDoPedro := contas.ContaCorrente{Titular: clientes.Titular{
		Nome: "Pedro Inácio",
		CPF: "54.345.323-55",
		Profissao: "Gerente"},
		NumeroAgencia: 12345,
		NumeroConta: 22,
		Saldo: 100}

	// Esboço
	contaDoJoao.Sacar(10)
	contaDoJoao.Depositar(10)

	// Transferência
	contaDoPedro.Transferir(50, &contaDoJoao)

	// Extrato
	fmt.Println(contaDoJoao.Titular.Nome, " - [ Saldo:", contaDoJoao.Saldo, "]")
	fmt.Println(contaDoPedro.Titular.Nome, " - [ Saldo:", contaDoPedro.Saldo, "]")
}