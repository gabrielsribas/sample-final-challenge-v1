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

func (s *ProductService) GetByID(id string) (model.Product, error) {
	return s.Repo.FindByID(id)
}

func (s *ProductService) GetByName(name string) ([]model.Product, error) {
	return s.Repo.FindByName(name)
}

func (s *ProductService) CountProducts() (int64, error) {
	return s.Repo.Count()
}

func (s *ProductService) CreateProduct(product model.Product) error {
	return s.Repo.Create(product)
}

func (s *ProductService) UpdateProduct(id string, product model.Product) error {
	return s.Repo.Update(id, product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.Repo.Delete(id)
}
