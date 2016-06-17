package main

import (
	"net/http"

	"github.com/brfreitas/restful-api/routers"
	"github.com/brfreitas/restful-api/settings"
	"github.com/codegangsta/negroni"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}
