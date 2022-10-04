package database

import (
	"RastreioBot/internal/pkg/models"
	"fmt"
	"log"
)

func (db Database) InsertPackage(userID string, name string, code string) (id int) {
	lastModified := models.GetLastModified(code)
	query := `INSERT INTO package (user_id, package_name, package_code, last_modified)
    VALUES ($1, $2, $3, $4)
    RETURNING id`

	err := db.database.QueryRow(query, userID, name, code, lastModified).Scan(&id)
	if err != nil {
		log.Printf("Ocorreu um erro ao inserir no banco de dados: %v\n", err)
	}

	return id
}

func (db Database) SelectByID(id int) (find bool, name string, code string) {
	query := `SELECT package_name, package_code FROM package WHERE id = $1`

	err := db.database.QueryRow(query, id).Scan(&name, &code)
	if err != nil {
		log.Printf("Erro ao selecionar o código pelo ID: %v\n", err)
		return false, "", ""
	}
	return true, name, code
}

func (db Database) SelectByName(name string) (find bool, code string) {
	query := `SELECT package_code FROM package WHERE package_name = $1`

	err := db.database.QueryRow(query, name).Scan(&code)
	if err != nil {
		log.Printf("Erro ao selecionar o código pelo ID: %v\n", err)
		return false, ""
	}
	return true, code
}

func (db Database) SelectByUserID(userID string) (find bool, result string) {
	query := `SELECT package_name, package_code FROM package WHERE user_id = $1`

	rows, err := db.database.Query(query, userID)
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

func (db Database) DeleteByName(name string) bool {
	query := `DELETE FROM package WHERE package_name = $1`

	_, err := db.database.Exec(query, name)
	if err != nil {
		log.Printf("Erro ao deletar um pacote: %v\n", err)
		return false
	}
	return true
}
