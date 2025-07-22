package models

import (
	"time"
	"gorm.io/gorm"
)

type Emprestimo struct {
	gorm.Model
	LivroID     uint      `gorm:"not null" json:"livro_id"`   
	UsuarioID   uint      `gorm:"not null" json:"usuario_id"`    
	DataInicio  time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"data_inicio"`
	DataFim     time.Time `gorm:"not null" json:"data_fim"`    
	Devolvido   bool      `gorm:"default:false" json:"devolvido"`

	Livro       Livro     `gorm:"foreignKey:LivroID" json:"livro"`     
	Usuario     Usuario   `gorm:"foreignKey:UsuarioID" json:"usuario"` 
}