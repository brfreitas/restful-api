package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brfreitas/restful-api/api/parameters"
	"github.com/brfreitas/restful-api/core/authentication"
	"github.com/brfreitas/restful-api/services/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Login metodo para realizar o login na api
func Login(requestUser *models.User) (int, []byte) {
	//Inicializa o backend do JWT
	authBackend := authentication.InitJWTAuthenticationBackend()
	//verifica login e senha
	if authBackend.Authenticate(requestUser) {
		//gera o token
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		}
		response, _ := json.Marshal(parameters.TokenAuthentication{token})
		return http.StatusOK, response
	}

	return http.StatusUnauthorized, []byte("")
}

//RefreshToken metodo para atualizar o token
func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

//Logout matodo para descartar o token
func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(req, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	fmt.Println(tokenString)
	fmt.Println(tokenRequest)
	return authBackend.Logout(tokenString, tokenRequest)
}
