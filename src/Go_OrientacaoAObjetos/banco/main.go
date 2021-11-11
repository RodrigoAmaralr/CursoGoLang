package main

import (
	"fmt"
)

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoRodrigo := contaCorrente{}
	contaDoRodrigo.titular = "Rodrigo"
	contaDoRodrigo.saldo = 500
	fmt.Println(contaDoRodrigo.saldo)

	fmt.Println(contaDoRodrigo.sacar(200))
	fmt.Println(contaDoRodrigo.saldo)

	status, valor := contaDoRodrigo.depositar(200)
	fmt.Println(status, valor)
	//println(strconv.FormatFloat(contaDoRodrigo.saldo, 'f', 2, 32))
}

func (c *contaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *contaCorrente) depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "O valor do deposito menor que zero", c.saldo
	}
}
