package routes

import (
	"agendamento-faculdade/controllers"

	"github.com/gin-gonic/gin"
)

func ProfissionalRoutes(router *gin.Engine) {
	profissionalGroup := router.Group("/profissionais")
	{
		profissionalGroup.POST("/", controllers.CriarProfissional)
		profissionalGroup.GET("/", controllers.ListarProfissionais)
		profissionalGroup.GET("/:id", controllers.BuscarProfissional)
		profissionalGroup.PUT("/:id", controllers.AtualizarProfissional)
		profissionalGroup.DELETE("/:id", controllers.DeletarProfissional)
	}
}
