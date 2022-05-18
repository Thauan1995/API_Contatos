package banco

import (
	"api/config"
	"api/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConectarPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.StringConexaoPostgres)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

//Insere dados na tabela contacts do Postgres
func InsereDadosPostgre(db *sql.DB, dados models.Cliente) error {
	for _, v := range dados.Contatos {
		sqlStatement := fmt.Sprintf("INSERT INTO contacts(nome, celular) VALUES('%s','%s');", v.Nome, v.Celular)
		res, err := db.Exec(sqlStatement)
		if err != nil {
			return fmt.Errorf("Falha na execução do Insert de contactos no MYSQL")
		}
		rowID, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		defer log.Printf("Inserida row: %d\n", rowID)
	}

	return nil
}
