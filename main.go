package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/GoTestApi/db"
	"github.com/GoTestApi/handlers"
)

func main() {
	router := gin.Default()
	// Conectar ao banco de dados
	db, err := db.AbrirConexao()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Configurar manipuladores de solicitações HTTP
	r := handlers.HandleRequest(db)

	// Iniciar servidor HTTP
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Erro ao iniciar servidor HTTP: %v", err)
	}
	//router.Run(":8080")
}
