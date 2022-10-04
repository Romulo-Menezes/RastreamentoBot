package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"os"
)

type config struct {
	Database string `json:"nameDB"`
	User     string `json:"userDB"`
	Password string `json:"passwordDB"`
	Host     string `json:"hostDB"`
	Port     string `json:"portDB"`
}

type Database struct {
	database *sql.DB
}

var sqlConn string

var (
	createTable = `
CREATE TABLE IF NOT EXISTS package(
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    package_name VARCHAR(255) NOT NULL,
    package_code VARCHAR (50) NOT NULL,
    last_modified VARCHAR (50)
);`
)

func init() {
	configJson, err := os.Open("config.json")
	if err != nil {
		log.Printf("Ocorreu um erro ao abrir o arquivo! %s\n", err.Error())
	}

	configBytes, err := ioutil.ReadAll(configJson)
	if err != nil {
		log.Printf("Ocorreu um erro ao converter para bytes! %s\n", err.Error())
	}

	var config config

	json.Unmarshal(configBytes, &config)

	sqlConn = fmt.Sprintf("host= %s port= %s user= %s password= %s dbname= %s",
		config.Host, config.Port, config.User, config.Password, config.Database)

	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Printf("Ocorreu um erro ao conectar no banco! %s\n", err.Error())
	}
	defer db.Close()

	_, err = db.Query(createTable)
	if err != nil {
		log.Printf("Ocorreu um erro ao tentar criar a tabela: %v\n", err)
	}
}

func New() Database {
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Printf("Ocorreu um erro ao conectar no banco! %s\n", err.Error())
	}
	return Database{database: db}
}

func (db Database) Close() {
	err := db.database.Close()
	if err != nil {
		log.Printf("Erro ao fechar a conexão com o banco: %v\n", err)
	}
}
