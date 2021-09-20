package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a url de comunicacao com a API
	APIURL = ""
	// Porta onde a aplicação web está rodando
	Porta = 0
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializar as variaveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HAS_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
}
