package main

import (
	"agendamento-faculdade/config"
	"agendamento-faculdade/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Conectar ao banco de dados
	config.ConnectDatabase()

	// Criar servidor Gin
	r := gin.Default()

	// Registrar rotas
	routes.AlunoRoutes(r)
	routes.ProfissionalRoutes(r)
	routes.AgendamentoRoutes(r)

	// Iniciar servidor
	port := "8080"
	fmt.Println("Servidor rodando na porta", port)
	r.Run(":" + port)
}