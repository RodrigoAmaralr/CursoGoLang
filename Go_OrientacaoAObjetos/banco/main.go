package main

import (
	"CursoGoLang/Go_OrientacaoAObjetos/banco/contas"
	"fmt"
)

/* Comandos necessario para incluir o contaCorrente.go
PS C:\Users\Rodrigo\go\src\CursoGoLang\Go_OrientacaoAObjetos\banco> go mod init CursoGoLang\Go_OrientacaoAObjetos\banco
PS C:\Users\Rodrigo\go\src\CursoGoLang\Go_OrientacaoAObjetos\banco>go mod tidy
*/

func main() {
	// contaDoRodrigo := contaCorrente{}
	// contaDoRodrigo.titular = "Rodrigo"
	// contaDoRodrigo.saldo = 500
	// fmt.Println(contaDoRodrigo.saldo)

	// fmt.Println(contaDoRodrigo.sacar(200))
	// fmt.Println(contaDoRodrigo.saldo)

	// status, valor := contaDoRodrigo.depositar(200)
	// fmt.Println(status, valor)

	///////////////////////////////
	// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 200}

	// status := contaDoGustavo.Transferir(200, &contaDaSilvia)
	// fmt.Println(status)

	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)

	////////////////////////////////
	// clienteBruno := clientes.Titular{Nome: "Bruno", CPF: "123.123.111.12", Profissao: "Desenvolvedor"}
	// contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}

	// fmt.Println(contaDoBruno)

	///////////////////////////////
	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(100)

	// fmt.Println(contaExemplo.ObterSaldo())

	////////////////
	// contaDoDenis := contas.ContaPoupanca{}
	// contaDoDenis.Depositar(100)
	// contaDoDenis.Sacar(500)

	// fmt.Println(contaDoDenis.ObterSaldo())

	////////////////////
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 400)

	fmt.Println(contaDaLuisa.ObterSaldo())

}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
