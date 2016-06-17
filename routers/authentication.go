package routers

import (
	"github.com/brfreitas/restful-api/controllers"
	"github.com/brfreitas/restful-api/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//SetAuthenticationRoutes define as rotas dos Endpoints de autenticacao
func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	//rota de login onde criam-se os tokens de acesso
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	//rota para atualizar o token
	router.Handle("/refresh-token-auth",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.RefreshToken),
		)).Methods("GET")
	//rota para invalidar o token
	router.Handle("/logout",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.Logout),
		)).Methods("GET")

	return router
}
