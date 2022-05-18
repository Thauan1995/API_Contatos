package tratativas

import (
	"api/banco"
	"api/models"
	"fmt"
	"regexp"
	"strings"
)

//Trata Especificações do cliente MACAPA
func ClienteMacapa(dadosClient *models.Cliente) error {
	var dadosFormat models.Cliente
	for _, dado := range dadosClient.Contatos {
		if dado.Nome != "" {
			dado.Nome = strings.ToUpper(dado.Nome)
		}

		if dado.Celular != "" {
			dado.Celular = AplicaRegexMacapa(dado.Celular)
		}
		dadosFormat.Contatos = append(dadosFormat.Contatos, dado)
	}

	db, err := banco.ConectarMYSQL()
	if err != nil {
		return fmt.Errorf("Falha ao conectar-se com o MYSQL")
	}
	err = banco.InserirDadosMY(db, dadosFormat)
	if err != nil {
		return fmt.Errorf("Erro ao inserir contatos no MYSQL")
	}
	return nil
}

//Trata Especificações do cliente VAREJAO
func ClienteVarejao(dadosClient *models.Cliente) error {
	db, err := banco.ConectarPostgres()
	if err != nil {
		return fmt.Errorf("Falha ao conectar-se com o Postgres")
	}
	err = banco.InsereDadosPostgre(db, *dadosClient)
	if err != nil {
		return fmt.Errorf("Erro ao inserir contatos no Postgres")
	}
	return nil
}

//Aplica Espressão regular pro cliente MACAPA
func AplicaRegexMacapa(celular string) string {
	var (
		brString     string
		dddString    string
		numeroString string
	)

	codeBR := regexp.MustCompile(".{2}")
	brString = codeBR.FindString(celular)

	codeDDD := regexp.MustCompile("[^5].")
	dddString = codeDDD.FindString(celular)

	regexNumero := fmt.Sprintf("[^%s%s](.*)", brString, dddString)

	codeNumero := regexp.MustCompile(regexNumero)
	numeroString = codeNumero.FindString(celular)

	celular = fmt.Sprintf("+%s (%s) %s", brString, dddString, numeroString)

	return celular

}
