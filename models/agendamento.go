package models

import (
	"time"
	"gorm.io/gorm"
)

type Agendamento struct {
	gorm.Model
	AlunoID       uint      `json:"aluno_id"`
	ProfissionalID uint      `json:"profissional_id"`
	DataHora      time.Time `json:"data_hora"`
	Status        string    `json:"status"`
}
