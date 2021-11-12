package main

// foi utilizado o postgresql-10.19-1 no computador
/*
Tools -> Query tools:
Query Editor
create table produtos (
	id serial primary key,
	nome varchar,
	descricao varchar,
	preco decimal,
	quantidade integer
)

insert into produtos (nome, descricao, preco, quantidade) values
('Camiseta', 'Preta', 19, 10),
('Fone', 'Muito bom', 99, 5);

insert into produtos (nome, descricao, preco, quantidade) values
('Produto novo', '2.0 mesmo', 199, 1);

*/

/*
go get github.com/lib/pq
go mod init
go mod tidy
*/

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	//conexao := "user dbname password host sslmode"
	conexao := "user=postgres dbname=alura_loja password=276813 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// db := conectaComBancoDeDados()
	// defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// produtos := []Produto{
	// 	{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
	// 	{"Tenis", "Confortável", 89, 3},
	// 	{"Fone", "Muito bom", 59, 2},
	// 	{"Produto novo", "Muito Legal", 1.99, 1},
	// }
	db := conectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
