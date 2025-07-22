package models

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model
	Nome      string      `gorm:"not null" json:"nome"`
	Email     string      `gorm:"unique;not null" json:"email"`
	
	Emprestimos []Emprestimo `gorm:"foreignKey:UsuarioID" json:"emprestimos"` 
}