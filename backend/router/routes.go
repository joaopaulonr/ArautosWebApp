package router


import (
	"backend/handler"
	
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// Rotas

func RegisterRoutes(router *gin.Engine, db *sqlx.DB, jwtKey []byte) {
	router.POST("/register", handler.HandleRegister(db, jwtKey))
	router.POST("/login", handler.HandleLogin(db, jwtKey))
}

func PostRoutes(router *gin.Engine, db *sqlx.DB, jwtKey []byte) {
	router.GET(("/post"), handler.ShowPost(db))
	router.POST(("/post"), handler.CreatePost(db))
	router.PUT(("/post"), handler.UpdatePost(db))
	router.DELETE(("/post"), handler.DeletePost(db))
	router.GET(("/posts"), handler.ListAllPosts(db))
}
