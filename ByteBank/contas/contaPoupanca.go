package contas

import (
	"fmt"
	"github.com/ygorsimoes/LearningGo/ByteBank/clientes"
)

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	Saldo         float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) {
	if valorDoSaque < 0 || valorDoSaque > c.Saldo {
		fmt.Println("[+] Não é possível sacar.")
	} else {
		c.Saldo -= valorDoSaque
		fmt.Println("[+] Saque no valor de R$", valorDoSaque, "realizado com sucesso!")
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) {
	if valorDoDeposito >= 0 {
		c.Saldo += valorDoDeposito
		fmt.Println("[+] Deposito no valor de", valorDoDeposito, "realizado com sucesso!")
	} else {
		fmt.Println("[+] Não é possível depositar. (Valor do deposito é menor que 0.)")
	}
}
