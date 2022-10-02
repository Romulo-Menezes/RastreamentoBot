package database

import (
	"RastreioBot/internal/pkg/models"
	"log"
)

func InsertPackage(userID string, name string, code string) (id int) {
	lastModified := models.GetLastModified(code)
	query := `INSERT INTO package (user_id, package_name, package_code, last_modified)
    VALUES ($1, $2, $3, $4)
    RETURNING id`
	db := DBConnect()

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

func SelectByID(id int) (find bool, name string, code string) {
	query := `SELECT package_name, package_code FROM package WHERE id = $1`
	db := DBConnect()
	defer db.Close()

	err := db.QueryRow(query, id).Scan(&name, &code)
	if err != nil {
		log.Printf("Erro ao selecionar o código pelo ID: %v\n", err)
		return false, "", ""
	}
	return true, name, code
}

func SelectByName(name string) (find bool, code string) {
	query := `SELECT package_code FROM package WHERE package_name = $1`
	db := DBConnect()
	defer db.Close()

	err := db.QueryRow(query, name).Scan(&code)
	if err != nil {
		log.Printf("Erro ao selecionar o código pelo ID: %v\n", err)
		return false, ""
	}
	return true, code
}

func DeleteByName(name string) bool {
	query := `DELETE FROM package WHERE package_name = $1`
	db := DBConnect()
	defer db.Close()

	_, err := db.Exec(query, name)
	if err != nil {
		log.Printf("Erro ao deletar um pacote: %v\n", err)
		return false
	}
	return true
}
