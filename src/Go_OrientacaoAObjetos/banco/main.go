package main

import (
	"fmt"
	//"Go_OrientacaoAObjetos/banco/contas"
)

//import "Go_OrientacaoAObjetos/banco/contas"

func main() {
	// contaDoRodrigo := contaCorrente{}
	// contaDoRodrigo.titular = "Rodrigo"
	// contaDoRodrigo.saldo = 500
	// fmt.Println(contaDoRodrigo.saldo)

	// fmt.Println(contaDoRodrigo.sacar(200))
	// fmt.Println(contaDoRodrigo.saldo)

	// status, valor := contaDoRodrigo.depositar(200)
	// fmt.Println(status, valor)

	contaDaSilvia := contas.contaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.contaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)
	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
