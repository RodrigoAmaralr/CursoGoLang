package main

import "fmt"

type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoRodrigo := contaCorrente{titular: "Rodrigo", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDoRabelo := contaCorrente{titular: "Rabelo", saldo: 50.9}
	contaDoAmaral := contaCorrente{"Amaral", 222, 111222, 200}
	fmt.Println(contaDoRodrigo)
	fmt.Println(contaDoRabelo)
	fmt.Println(contaDoAmaral)
}
