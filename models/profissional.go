package models

import "gorm.io/gorm"

type Profissional struct {
	gorm.Model
	Nome        string `json:"nome"`
	Email       string `json:"email" gorm:"unique"`
	Especialidade string `json:"especialidade"`
	Telefone    string `json:"telefone"`
}
