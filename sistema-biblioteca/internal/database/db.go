package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("biblioteca.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    DB = db
    return db, nil
}