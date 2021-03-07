package contas

import (
	"fmt"
	"github.com/ygorsimoes/LearningGo/ByteBank/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) {
	if valorDoSaque < 0 || valorDoSaque > c.Saldo {
		fmt.Println("[+] Não é possível sacar.")
	} else {
		c.Saldo -= valorDoSaque
		fmt.Println("[+] Saque no valor de R$", valorDoSaque, "realizado com sucesso!")
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) {
	if valorDoDeposito >= 0 {
		c.Saldo += valorDoDeposito
		fmt.Println("[+] Deposito no valor de", valorDoDeposito, "realizado com sucesso!")
	} else {
		fmt.Println("[+] Não é possível depositar. (Valor do deposito é menor que 0.)")
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente)  {
	if valorDaTransferencia < 0 || valorDaTransferencia > c.Saldo {
		fmt.Println("[+] Não é possível transferir.")
	} else {
		c.Saldo -= valorDaTransferencia
		contaDestino.Saldo += valorDaTransferencia
		fmt.Println("[+] Transferência no valor de", valorDaTransferencia, "realizado com sucesso!")
	}
}
