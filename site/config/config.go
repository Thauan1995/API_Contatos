package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoMYSQL    = ""
	StringConexaoPostgres = ""
	Port                  = 0
	SecretKey             []byte
)

// Carregar() vai inicializar as variáveis de ambiente
func Carregar() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	StringConexaoMYSQL = fmt.Sprintf("%s:%s@tcp(localhost:13306)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"),
	)

	StringConexaoPostgres = fmt.Sprintf("host=localhost port=15432 user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
