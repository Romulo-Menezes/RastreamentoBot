package database

import (
	"RastreioBot/internal/pkg/models"
	"log"
)

func InsertPackage(userID string, name string, code string) int {
	lastModified := models.GetLastModified(code)
	query := `INSERT INTO package (user_id, package_name, package_code, last_modified)
    VALUES ($1, $2, $3, $4)
    RETURNING id`
	db := DBConnect()
	id := 0

	err := db.QueryRow(query, userID, name, code, lastModified).Scan(&id)
	if err != nil {
		log.Printf("Ocorreu um erro ao inserir no banco de dados: %v\n", err)
	}

	err = db.Close()
	if err != nil {
		log.Printf("Ocorreu um erro ao fechar a conexão com o banco de dados: %v\n", err)
	}
	return id
}
