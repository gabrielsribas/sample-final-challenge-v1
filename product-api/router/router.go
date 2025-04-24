package router

import (
	"product-api/controller"
	"product-api/model"
	"product-api/repository"
	"product-api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	db.AutoMigrate(&model.Product{})

	repo := &repository.ProductRepository{DB: db}
	service := &service.ProductService{Repo: repo}
	controller := &controller.ProductController{Service: service}

	r.GET("/products", controller.GetAll)
	r.GET("/products/:id", controller.GetByID)
	r.GET("/products/name/:name", controller.GetByName)
	r.GET("/products/count", controller.Count)
	r.POST("/products", controller.Create)
	r.PUT("/products/:id", controller.Update)
	r.DELETE("/products/:id", controller.Delete)
}
