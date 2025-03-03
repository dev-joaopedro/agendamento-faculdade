package routes

import (
	"agendamento-faculdade/controllers"

	"github.com/gin-gonic/gin"
)

func AlunoRoutes(router *gin.Engine) {
	alunoGroup := router.Group("/alunos")
	{
		alunoGroup.POST("/", controllers.CriarAluno)
		alunoGroup.GET("/", controllers.ListarAlunos)
		alunoGroup.GET("/:id", controllers.BuscarAluno)
		alunoGroup.PUT("/:id", controllers.AtualizarAluno)
		alunoGroup.DELETE("/:id", controllers.DeletarAluno)
	}
}
