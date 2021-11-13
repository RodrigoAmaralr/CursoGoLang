package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	//conexao := "user dbname password host sslmode"
	//conexao := "user=postgres dbname=alura_loja password=276813 host=localhost sslmode=disable port=5432"	//Windows
	conexao := "user=postgres dbname=alura_loja password=276813 host=localhost sslmode=disable port=5433" //Linux
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
