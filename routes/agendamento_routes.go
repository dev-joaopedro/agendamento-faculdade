package routes

import (
	"agendamento-faculdade/controllers"

	"github.com/gin-gonic/gin"
)

func AgendamentoRoutes(router *gin.Engine) {
	agendamentoGroup := router.Group("/agendamentos")
	{
		agendamentoGroup.POST("/", controllers.CriarAgendamento)
		agendamentoGroup.GET("/", controllers.ListarAgendamentos)
		agendamentoGroup.GET("/:id", controllers.BuscarAgendamento)
		agendamentoGroup.PUT("/:id", controllers.AtualizarAgendamento)
		agendamentoGroup.DELETE("/:id", controllers.DeletarAgendamento)
	}
}
