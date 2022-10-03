package database

import (
	"RastreioBot/internal/pkg/models"
	"fmt"
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

func SelectByUserID(userID string) (find bool, result string) {
	query := `SELECT package_name, package_code FROM package WHERE user_id = $1`
	db := DBConnect()
	defer db.Close()

	rows, err := db.Query(query, userID)
	if err != nil {
		log.Printf("Erro ao selecionar pelo ID do usuário: %v\n", err)
		return false, ""
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var code string

		err = rows.Scan(&name, &code)
		if err != nil {
			log.Printf("Erro ao ler a linha da tabela: %v\n", err)
			return false, ""
		}
		result = fmt.Sprintf("%v\n%v - %v\n", result, name, code)
	}
	return true, result
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
