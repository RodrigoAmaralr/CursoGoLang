package main

import "fmt"

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
