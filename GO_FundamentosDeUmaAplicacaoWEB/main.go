package main

import (
	"CursoGoLang/GO_FundamentosDeUmaAplicacaoWEB/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
