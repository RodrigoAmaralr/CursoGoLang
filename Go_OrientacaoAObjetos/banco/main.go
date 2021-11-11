package main

/* Comandos necessario para incluir o contaCorrente.go
PS C:\Users\Rodrigo\go\src\CursoGoLang\Go_OrientacaoAObjetos\banco> go mod init CursoGoLang\Go_OrientacaoAObjetos\banco
PS C:\Users\Rodrigo\go\src\CursoGoLang\Go_OrientacaoAObjetos\banco>go mod tidy
*/

import (
	"CursoGoLang/Go_OrientacaoAObjetos/banco/contas"
	"fmt"
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

	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDoGustavo.Transferir(-200, &contaDaSilvia)
	fmt.Println(status)

	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
