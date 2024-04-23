package rotas

import (
	"api/src/controllers"
	"net/http"
)

const URI_USUARIOS = "/usuarios"
const URI_USUARIOS_BY_ID = "/usuarios/{usuarioId}"

var rotasUsuarios = []Rota{
	{
		URI:                URI_USUARIOS,
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                URI_USUARIOS,
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarios,
		RequerAutenticacao: false,
	},
	{
		URI:                URI_USUARIOS_BY_ID,
		Metodo:             http.MethodPut,
		Funcao:             controllers.AlterarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                URI_USUARIOS_BY_ID,
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                URI_USUARIOS_BY_ID,
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: false,
	},
}
