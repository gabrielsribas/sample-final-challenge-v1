package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
    if err != nil {
        panic("Erro ao conectar com banco de dados")
    }
    return db
}

