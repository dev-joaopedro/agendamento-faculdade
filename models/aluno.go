package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	Nome   string `json:"nome"`
	Email  string `json:"email" gorm:"unique"`
	Curso  string `json:"curso"`
	Telefone string `json:"telefone"`
}
