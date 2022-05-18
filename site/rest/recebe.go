package rest

import (
	"api/autenticacao"
	"api/banco"
	"api/models"
	"api/tratativas"
	"api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func RecebeDadosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		RecebeDados(w, r)
		return
	}
	utils.RespondWithError(w, http.StatusBadRequest, 0, "Método não permitido")
	return
}

func RecebeDados(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao receber body do cliente")
		return
	}

	if err := json.Unmarshal(body, &cliente); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal do cliente")
		return
	}

	usuarioIDToken, err := autenticacao.ExtrairUsuarioID(r)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao extrair id do usuario do token")
		return
	}

	db, err := banco.ConectarMYSQL()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Falha ao concectar com o mysql")
		return
	}
	defer db.Close()

	usuarioBanco, err := banco.BuscarUsuarioPorID(db, usuarioIDToken)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao buscar usuario por ID")
		return
	}

	switch strings.ToUpper(usuarioBanco.Nome) {
	case "MACAPA":
		err := tratativas.ClienteMacapa(&cliente)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro nas tratativas do cliente Macapá")
			return
		}
	case "VAREJAO":
		err := tratativas.ClienteVarejao(&cliente)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro nas tratativas do cliente Macapá")
			return
		}
	default:
		utils.RespondWithError(w, http.StatusNotFound, 0, "Nenhuma tratativa encontrada para o cliente "+usuarioBanco.Nome)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Dados inseridos com sucesso!")
}
