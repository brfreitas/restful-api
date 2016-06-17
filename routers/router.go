package routers

import (
	"github.com/gorilla/mux"
)

//InitRoutes Inicializa todas as rotas da API
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
