// internal/database/db.go
package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/llopes05/RESTeSOAP/internal/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("biblioteca.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Livro{},
		&models.Usuario{},
		&models.Emprestimo{},
	)
	DB = db
	return db, err
}