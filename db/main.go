package db

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type ConfigDb struct {
	Host 			string
	Port 			int
	User			string
	Password 		string
	NameDataBase 	string
}

// Variável que guardará a conexão com o banco postgres
var conn *sql.DB

// Função que estabelecerá a conexão com o banco de dados
func ConnectDb(config ConfigDb) (*sql.DB, error)  {
	var err error
	
	// Aqui fazemos a abertura de conexão com banco postegres
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.NameDataBase)

	conn, err = sql.Open("postgres", connStr)

	// Condição para tratar algum erro caso aconteça
	if (err != nil) {
		return nil, err
	}

    return conn, nil
}

// Função que fecha a conexão com o banco de dados
func CloseDb()  {
	conn.Close()
}
