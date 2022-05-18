package middlewares

import (
	"api/autenticacao"
	"api/utils"
	"net/http"
)

// Autenticar verifica se o usuario fazendo a requisição está autenticado
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, 0, "Erro ao validar Token")
			return
		}
		proximaFuncao(w, r)
	}
}
