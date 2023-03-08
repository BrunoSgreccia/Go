package handlers

import (
	"database/sql"
	"net/http"

	"github.com/GoTestApi/db"
	"github.com/gin-gonic/gin"
)

// HandleRequest configura os manipuladores de solicitações HTTP
func HandleRequest(dsqlb *sql.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/usuarios", func(c *gin.Context) {
		usuarios, err := db.TodosUsuarios(dsqlb)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao buscar usuários",
			})
			return
		}

		c.JSON(http.StatusOK, usuarios)
	})

	r.POST("/usuarios", func(c *gin.Context) {
		var usuario db.Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao analisar entrada",
			})
			return
		}

		if err := db.NovoUsuario(dsqlb, usuario.Nome, usuario.Email, usuario.Telefone); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Erro ao adicionar usuário",
			})
			return
		}

		c.Status(http.StatusCreated)
	})

	return r
}
