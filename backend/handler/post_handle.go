package handler

import (
	"backend/middleware"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func ShowPost(db *sqlx.DB, jwtKey []byte) {

}
func CreatePost(db *sqlx.DB, jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obter o ID do usuário do token usando o middleware
		userID, err := middleware.GetUserIDFromToken(c, jwtKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Agora você tem o userID, que você pode usar na criação da postagem
		var post models.Post
		if err := c.BindJSON(&post); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		// Definir o autor da postagem como o usuário atual
		post.AuthorID = userID

		// Inserir a postagem no banco de dados
		_, err = db.NamedExec("INSERT INTO posts (author_id, title, content) VALUES (:author_id, :title, :content)", post)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating post"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
	}
}
func UpdatePost(db *sqlx.DB, jwtKey []byte) {

}
func DeletePost(db *sqlx.DB, jwtKey []byte) {

}
func ListAllPosts(db *sqlx.DB, jwtKey []byte) {

}
