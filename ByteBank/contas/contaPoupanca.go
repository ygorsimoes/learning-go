package contas

import (
	"github.com/ygorsimoes/LearningGo/ByteBank/clientes"
	"strconv"
)

type ContaPoupanca struct {
	titular                              clientes.Titular
	numeroAgencia, numeroConta, Operacao int
	saldo                                float64
}

func (c ContaPoupanca) Sacar(valorDoSaque float64) string {
	if valorDoSaque > 0 && valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado realizado com sucesso!"
	} else {
		return "Saldo insuficiente!"
	}
}

func (c ContaPoupanca) Depositar(valorDoDeposito float64) string {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso!"
	} else {
		return "Deposito inválido!"
	}
}

func (c ContaPoupanca) Extrato() string {
	extrato := "=====================================" +
		"\nTitular: " + c.titular.Nome +
		"\nAgência: " + strconv.Itoa(c.numeroAgencia) +
		"\nConta: " + strconv.Itoa(c.numeroConta) +
		"\nsaldo: " + strconv.Itoa(int(c.saldo)) +
		"\n====================================="

	return extrato
}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}
