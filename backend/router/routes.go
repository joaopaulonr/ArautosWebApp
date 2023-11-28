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
	// router.GET(("/post"), handler.ShowPost(db, jwtKey))
	router.POST(("/post"), handler.CreatePost(db, jwtKey))
	// router.PUT(("/post"), handler.UpdatePost(db, jwtKey))
	// router.DELETE(("/post"), handler.DeletePost(db,jwtKey))
	// router.GET(("/posts"), handler.ListAllPosts(db, jwtKey))
}
