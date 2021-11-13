package routes

import (
	"CursoGoLang/GO_FundamentosDeUmaAplicacaoWEB/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
