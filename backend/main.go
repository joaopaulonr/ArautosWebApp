package main

import (
	"backend/db"
	"backend/router"
	"fmt"
)

func main() {
	if err := db.InitializeDB(); err != nil {
		fmt.Println("Erro ao iniciar o banco de dados:", err)
		return
	}
	router.Initialize()
}
