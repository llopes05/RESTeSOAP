package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/llopes05/RESTeSOAP/internal/models"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	if DB != nil {
		return DB, nil 
	}
	db, err := gorm.Open(sqlite.Open("biblioteca.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Livro{},
		&models.Emprestimo{},
		&models.Usuario{}, 
	)
	if err != nil {
		return nil, err
	}

	DB = db
	return DB, nil
}