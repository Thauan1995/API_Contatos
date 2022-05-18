package banco

import (
	"api/config"
	"api/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//Abre conexão com MYSQL e a retorna
func ConectarMYSQL() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConexaoMYSQL)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Busca o usuario cadastrado no banco pelo nome
func BuscarUsuarioPorNome(db *sql.DB, nome string) (models.Usuario, error) {
	linha, err := db.Query("select id, senha from usuarios where nome = ?", nome)
	if err != nil {
		return models.Usuario{}, fmt.Errorf("Erro na query para buscar usuario por nome")
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Senha); err != nil {
			return models.Usuario{}, fmt.Errorf("Erro na iteração das rows na struct usuario")
		}
	}

	return usuario, nil
}

//Buscar o usuario cadastrado no banco pelo ID
func BuscarUsuarioPorID(db *sql.DB, id int64) (models.Usuario, error) {
	linha, err := db.Query("select id, nome from usuarios where id = ?", id)
	if err != nil {
		return models.Usuario{}, fmt.Errorf("Erro na query para buscar usuario por id")
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if err = linha.Scan(&usuario.ID, &usuario.Nome); err != nil {
			return models.Usuario{}, fmt.Errorf("Erro na iteração das rows na struct usuario")
		}
	}

	return usuario, nil
}

//Insere dados na tabela contacts do MYSQL
func InserirDadosMY(db *sql.DB, dados models.Cliente) error {
	for _, v := range dados.Contatos {
		sqlStatement := fmt.Sprintf("INSERT INTO contacts(nome, celular) VALUES('%s','%s');", v.Nome, v.Celular)
		res, err := db.Exec(sqlStatement)
		if err != nil {
			return fmt.Errorf("Falha na execução do Insert de contactos no MYSQL")
		}
		rowID, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		defer log.Printf("Inserida row: %d\n", rowID)
	}

	return nil
}
