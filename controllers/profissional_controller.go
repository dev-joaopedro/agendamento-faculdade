package controllers

import (
	"agendamento-faculdade/config"
	"agendamento-faculdade/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Criar um novo profissional
func CriarProfissional(c *gin.Context) {
	var profissional models.Profissional
	if err := c.ShouldBindJSON(&profissional); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&profissional)
	c.JSON(http.StatusCreated, profissional)
}

// Listar todos os profissionais
func ListarProfissionais(c *gin.Context) {
	var profissionais []models.Profissional
	config.DB.Find(&profissionais)
	c.JSON(http.StatusOK, profissionais)
}

// Buscar profissional por ID
func BuscarProfissional(c *gin.Context) {
	var profissional models.Profissional
	id := c.Param("id")

	if err := config.DB.First(&profissional, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profissional não encontrado"})
		return
	}

	c.JSON(http.StatusOK, profissional)
}

// Atualizar profissional
func AtualizarProfissional(c *gin.Context) {
	var profissional models.Profissional
	id := c.Param("id")

	if err := config.DB.First(&profissional, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profissional não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&profissional); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&profissional)
	c.JSON(http.StatusOK, profissional)
}

// Deletar profissional
func DeletarProfissional(c *gin.Context) {
	var profissional models.Profissional
	id := c.Param("id")

	if err := config.DB.First(&profissional, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profissional não encontrado"})
		return
	}

	config.DB.Delete(&profissional)
	c.JSON(http.StatusOK, gin.H{"message": "Profissional deletado com sucesso"})
}
