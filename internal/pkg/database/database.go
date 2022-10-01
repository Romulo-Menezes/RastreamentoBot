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

type database struct {
	Database string `json:"nameDB"`
	User     string `json:"userDB"`
	Password string `json:"passwordDB"`
	Host     string `json:"hostDB"`
	Port     string `json:"portDB"`
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

	var config database

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

func DBConnect() *sql.DB {
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Printf("Ocorreu um erro ao conectar no banco! %s\n", err.Error())
	}
	return db
}
