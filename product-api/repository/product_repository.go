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

func (r *ProductRepository) FindByID(id string) (model.Product, error) {
	var product model.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) FindByName(name string) ([]model.Product, error) {
	var products []model.Product
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
}

func (r *ProductRepository) Count() (int64, error) {
	var count int64
	err := r.DB.Model(&model.Product{}).Count(&count).Error
	return count, err
}

func (r *ProductRepository) Create(product model.Product) error {
	return r.DB.Create(&product).Error
}

func (r *ProductRepository) Update(id string, newProduct model.Product) error {
	var product model.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		return err
	}
	product.Name = newProduct.Name
	product.Price = newProduct.Price
	return r.DB.Save(&product).Error
}

func (r *ProductRepository) Delete(id string) error {
	return r.DB.Delete(&model.Product{}, id).Error
}
