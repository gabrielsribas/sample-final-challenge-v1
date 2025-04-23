package service

import (
    "product-api/model"
    "product-api/repository"
)

type ProductService struct {
    Repo *repository.ProductRepository
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
    return s.Repo.FindAll()
}

// ... outros métodos: GetByID, Create, Update, Delete, Count, FindByName

