// router/router.go
package router

import (
	"backend/config"
	"backend/db"
	"backend/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	// Carregar a chave JWT do arquivo .env usando o pacote config
	jwtKey := config.LoadConfig()

	// Configurar o middleware de autenticação com a chave JWT
	router.Use(middleware.AuthMiddleware(jwtKey))

	// Registrar rotas de autenticação
	RegisterRoutes(router, db.DB, jwtKey)
	
	// Rotas das postagens
	PostRoutes(router,db.DB, jwtKey)

	// Iniciar o servidor
	if err := router.Run(":5000"); err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
