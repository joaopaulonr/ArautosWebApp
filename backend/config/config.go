// config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig carrega a configuração do arquivo .env
func LoadConfig() []byte {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar variáveis de ambiente:", err)
	}

	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		log.Fatal("A variável JWT_KEY não está definida no arquivo .env")
	}

	return []byte(jwtKey)
}
