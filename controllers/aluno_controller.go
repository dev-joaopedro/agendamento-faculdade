package controllers

import (
	"agendamento-faculdade/config"
	"agendamento-faculdade/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Criar um novo aluno
func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&aluno)
	c.JSON(http.StatusCreated, aluno)
}

// Listar todos os alunos
func ListarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	config.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

// Buscar aluno por ID
func BuscarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// Atualizar aluno
func AtualizarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// Deletar aluno
func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Param("id")

	if err := config.DB.First(&aluno, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Aluno não encontrado"})
		return
	}

	config.DB.Delete(&aluno)
	c.JSON(http.StatusOK, gin.H{"message": "Aluno deletado com sucesso"})
}
