//Package settings Configura os diferentes ambientes de execucao da API (preproduction, tests, production)
package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/pre.json",
	"tests":         "../../settings/tests.json",
}

//Settings definição da estrutura de configuracao do ambiente
type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

//Settings instancia do ambiente de execucao atual
var settings = Settings{}

var env = "preproduction"

//Init incializa as configuracoes de acordo com o ambiente
func Init() {
	//tenta pegar o ambiente na variavel de ambientes
	env = os.Getenv("GO_ENV")
	if env == "" {
		fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
		env = "preproduction"
	}
	LoadSettingsByEnv(env)
}

//LoadSettingsByEnv carrega o json de acordo com o ambiente
func LoadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error while reading config file", err)
	}
	settings = Settings{}
	jsonErr := json.Unmarshal(content, &settings)
	if jsonErr != nil {
		fmt.Println("Error while parsing config file", jsonErr)
	}
}

//GetEnvironment retorna o ambiente de execucao atual
func GetEnvironment() string {
	return env
}

//Get retorna as configuracoes atuais
func Get() Settings {
	if &settings == nil {
		Init()
	}
	return settings
}

//IsTestEnvironment verifica se e ambiente de teste
func IsTestEnvironment() bool {
	return env == "tests"
}
