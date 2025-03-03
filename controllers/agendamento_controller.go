package controllers

import (
	"agendamento-faculdade/config"
	"agendamento-faculdade/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Criar um novo agendamento
func CriarAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&agendamento)
	c.JSON(http.StatusCreated, agendamento)
}

// Listar todos os agendamentos
func ListarAgendamentos(c *gin.Context) {
	var agendamentos []models.Agendamento
	config.DB.Find(&agendamentos)
	c.JSON(http.StatusOK, agendamentos)
}

// Buscar agendamento por ID
func BuscarAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	id := c.Param("id")

	if err := config.DB.First(&agendamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	c.JSON(http.StatusOK, agendamento)
}

// Atualizar agendamento
func AtualizarAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	id := c.Param("id")

	if err := config.DB.First(&agendamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&agendamento); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&agendamento)
	c.JSON(http.StatusOK, agendamento)
}

// Deletar agendamento
func DeletarAgendamento(c *gin.Context) {
	var agendamento models.Agendamento
	id := c.Param("id")

	if err := config.DB.First(&agendamento, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agendamento não encontrado"})
		return
	}

	config.DB.Delete(&agendamento)
	c.JSON(http.StatusOK, gin.H{"message": "Agendamento deletado com sucesso"})
}
