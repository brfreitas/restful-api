package routers

import (
	"github.com/brfreitas/restful-api/controllers"
	"github.com/brfreitas/restful-api/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

//SetHelloRoutes define as rotas dos endpoints hello
func SetHelloRoutes(router *mux.Router) *mux.Router {
	//adiciona um midlewere para verificar a autenticacao do usuario
	router.Handle("/test/hello",
		negroni.New(negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController))).Methods("GET")
	return router
}
