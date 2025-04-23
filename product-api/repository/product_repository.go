package repository

import (
    "product-api/model"
    "gorm.io/gorm"
)

type ProductRepository struct {
    DB *gorm.DB
}

func (r *ProductRepository) FindAll() ([]model.Product, error) {
    var products []model.Product
    err := r.DB.Find(&products).Error
    return products, err
}

// ... outros métodos: FindByID, Create, Update, Delete

