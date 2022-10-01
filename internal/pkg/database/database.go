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

	sqlConn = fmt.Sprintf("host= %s user= %s password= %s dbname= %s",
		config.Host, config.User, config.Password, config.Database)
}

func DBConnect() *sql.DB {
	db, err := sql.Open("postgres", sqlConn)
	if err != nil {
		log.Printf("Ocorreu um erro ao conectar no banco! %s\n", err.Error())
	}
	return db
}
