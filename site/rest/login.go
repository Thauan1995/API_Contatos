package rest

import (
	"api/autenticacao"
	"api/banco"
	"api/models"
	"api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		Login(w, r)
		return
	}

	utils.RespondWithError(w, http.StatusMethodNotAllowed, 0, "Método não permitido")
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Erro ao receber body da requisição")
		return
	}

	var usuario models.Usuario
	if err = json.Unmarshal(body, &usuario); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, 0, "Falha ao realizar unmarshal da requisição")
		return
	}

	db, err := banco.ConectarMYSQL()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Falha ao concectar com o mysql")
		return
	}
	defer db.Close()

	usuarioBanco, err := banco.BuscarUsuarioPorNome(db, usuario.Nome)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, 0, "Falha ao buscar usuario")
		return
	}

	if usuario.Senha != usuarioBanco.Senha {
		utils.RespondWithError(w, http.StatusUnauthorized, 0, "Senha incorreta")
		return
	}

	token, _ := autenticacao.CriarToken(usuarioBanco.ID)
	utils.RespondWithJSON(w, http.StatusOK, token)

}
