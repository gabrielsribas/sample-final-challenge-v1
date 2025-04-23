package router

import (
    "product-api/controller"
    "product-api/repository"
    "product-api/service"
    "product-api/model"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
    db.AutoMigrate(&model.Product{})

    repo := &repository.ProductRepository{DB: db}
    srv := &service.ProductService{Repo: repo}
    ctrl := &controller.ProductController{Service: srv}

    r.GET("/products", ctrl.GetAll)
    // outros endpoints: GET by ID, POST, PUT, DELETE, COUNT, FIND BY NAME
}

