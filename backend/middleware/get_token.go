package middleware

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetUserIDFromToken(c *gin.Context, jwtKey []byte) (int64, error) {
	// Obtém o token da solicitação
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return 0, errors.New("no token provided")
	}

	// Remove o prefixo "Bearer " do token
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Faz a validação do token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	// Verifica se as claims são do tipo esperado
	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	// Obtém o ID do usuário das claims
	userID := claims.ID
	return userID, nil
}

// CustomClaims é uma estrutura que representa as claims do token
type CustomClaims struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}
