package controllers

import (
	"net/http"
)

//HelloController define a saida para o endpoint
func HelloController(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Write([]byte("Hello, World!"))
}
