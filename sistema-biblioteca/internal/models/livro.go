package models

import "gorm.io/gorm"

type Livro struct {
    gorm.Model
    Titulo     string `gorm:"not null" json:"titulo"`
    Autor      string `gorm:"not null" json:"autor"`
    Ano        int    `gorm:"not null" json:"ano"`
    Disponivel bool   `gorm:"default:true" json:"disponivel"`

    Emprestimos []Emprestimo `gorm:"foreignKey:LivroID" json:"emprestimos"`
}